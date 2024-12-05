package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/urfave/cli/v2"

	"github.com/education-english-web/BE-english-web/app/config"
	"github.com/education-english-web/BE-english-web/app/external/framework"
	pkgConfig "github.com/education-english-web/BE-english-web/pkg/config"
	"github.com/education-english-web/BE-english-web/pkg/log"
	"github.com/education-english-web/BE-english-web/pkg/middleware"
	"github.com/education-english-web/BE-english-web/pkg/profiler"
	"github.com/education-english-web/BE-english-web/pkg/redis"
	"github.com/education-english-web/BE-english-web/pkg/tracer"
)

const (
	// Constant for notifier engine
	notifierEngineRollbar = "rollbar"
	logLevelDebug         = "debug"
)

type service struct {
	cfg *config.Config
}

func newService(ctx *cli.Context) *service {
	s := &service{}

	s.loadConfig(ctx)

	if err := log.SetLevel(s.cfg.LogLevel); err != nil {
		panic(err)
	}

	log.AddHook(s.getLogHook(s.cfg.NotifierEngine))
	tracer.SetTracer(s.getTracer(s.cfg.TracerEngine))

	return s
}

func (s *service) loadConfig(ctx *cli.Context) {
	conf := &config.Config{
		Env: ctx.String(EnvFlag.Name),
		HTTPServer: config.ServerAddr{
			Port: ctx.String(HTTPPortFlag.Name),
		},
		Postgres: config.Postgres{
			ConnectionString: ctx.String(PostgresConnFlag.Name),
			Masters:          ctx.String(PostgresMasterHostsFlag.Name),
			Slaves:           ctx.String(PostgresSlaveHostsFlag.Name),
			Port:             ctx.Int(PostgresPortFlag.Name),
			User:             ctx.String(PostgresUserFlag.Name),
			Password:         ctx.String(PostgresPasswordFlag.Name),
			DB:               ctx.String(PostgresDatabaseFlag.Name),
			MaxOpenConns:     ctx.Int(PostgresMaxOpenConnsFlag.Name),
			MaxIdleConns:     ctx.Int(PostgresMaxIdleConnsFlag.Name),
			ConnMaxLifetime:  ctx.Int(PostgresConnMaxLifetimeFlag.Name),
			IsEnabledLog:     ctx.String(LogLevelFlag.Name) == logLevelDebug,
		},
		Redis: redis.Config{
			ConnectionString:   ctx.String(RedisConnFlag.Name),
			Host:               ctx.String(RedisHostFlag.Name),
			Port:               ctx.Int(RedisPortFlag.Name),
			User:               ctx.String(RedisUserFlag.Name),
			Pass:               ctx.String(RedisPasswordFlag.Name),
			Database:           ctx.Int(RedisDatabaseFlag.Name),
			PoolSize:           ctx.Int(RedisPoolSizeFlag.Name),
			InsecureSkipVerify: ctx.Bool(RedisInsecureSkipVerifyFlag.Name),
			TLS:                ctx.Bool(RedisEnabledTLSFlag.Name),
		},
		Datadog: config.Datadog{
			Env:            ctx.String(EnvFlag.Name),
			ServiceName:    ctx.String(AppNameFlag.Name),
			ServiceVersion: ctx.String(AppVersionFlag.Name),
			Host:           ctx.String(DatadogAgentHostFlag.Name),
			AgentAMPPort:   ctx.String(DatadogAgentAPMPortFlag.Name),
		},
		CORS: config.CORS{
			AllowHosts: strings.Split(ctx.String(CORSAllowHostsFlag.Name), ","),
		},
		Rollbar: config.Rollbar{
			Token: ctx.String(RollbarTokenFlag.Name),
			Env:   ctx.String(EnvFlag.Name),
		},
		GSuite: config.GSuite{
			ClientID:     ctx.String(GSuiteClientIDFlag.Name),
			ClientSecret: ctx.String(GSuiteClientSecretFlag.Name),
		},

		GServiceAccount: config.GServiceAccount{
			Email:                            ctx.String(GServiceAccountEmailFlag.Name),
			PrivateKey:                       ctx.String(GServiceAccountPrivateKeyFlag.Name),
			PrivateKeyID:                     ctx.String(GServiceAccountPrivateKeyIDFlag.Name),
			SpreadSheetIDOfficesReport:       ctx.String(GServiceSpreadSheetIDOfficesReportFlag.Name),
			SpreadSheetIDOfficesStatusReport: ctx.String(GServiceSpreadSheetIDOfficesStatusReportFlag.Name),
		},
		EnabledProfiling: ctx.Bool(EnabledProfilingFlag.Name),
		ProfilerEngine:   ctx.String(ProfilerEngineFlag.Name),
		LogLevel:         ctx.String(LogLevelFlag.Name),
		TracerEngine:     ctx.String(TracerEngineFlag.Name),
		ProxyURL:         ctx.String(ProxyURLFlag.Name),
		OIDCAllowDomains: strings.Split(ctx.String(OIDCAllowDomains.Name), ","),
		Salt:             ctx.String(SaltFlag.Name),
		JWTSecret:        ctx.String(JWTSecretFlag.Name),
	}
	s.cfg = conf

	pkgConfig.SetConfig(conf)
}

func (s *service) getLogHook(engine string) log.Hook {
	switch engine {
	case "":
		return log.NewNopHook()
	case notifierEngineRollbar:
		return log.NewRollbarHook(s.cfg.Rollbar.Token, s.cfg.Rollbar.Env)
	default:
		panic(fmt.Errorf("unsupported log hook engine: %s", engine))
	}
}

func (s *service) getTracer(engine string) tracer.Tracer {
	switch engine {
	case "":
		return tracer.NopTracer{}
	case middleware.TracerEngineNop:
		return tracer.NopTracer{}
	case middleware.TracerEngineDatadog:
		return tracer.NewDatadogTracer(&tracer.DatadogConfig{
			Env:            s.cfg.Datadog.Env,
			ServiceName:    s.cfg.Datadog.ServiceName,
			ServiceVersion: s.cfg.Datadog.ServiceVersion,
			Host:           s.cfg.Datadog.Host,
			AgentAMPPort:   s.cfg.Datadog.AgentAMPPort,
		})
	default:
		panic(fmt.Errorf("unsupported tracer engine: %s", engine))
	}
}

func (s *service) getProfiler(engine string) profiler.Profiler {
	switch engine {
	case "":
		return profiler.NopProfiler{}
	case middleware.ProfilerEngineNop:
		return profiler.NopProfiler{}
	case middleware.ProfilerEngineDatadog:
		return profiler.NewDatadogProfiler(&profiler.DatadogConfig{
			Env:            s.cfg.Datadog.Env,
			ServiceName:    s.cfg.Datadog.ServiceName,
			ServiceVersion: s.cfg.Datadog.ServiceVersion,
		})
	default:
		panic(fmt.Errorf("unsupported profiler engine: %s", engine))
	}
}

func (s *service) start() error {
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM)

	s.initDBConnectionPostgres(s.cfg.Postgres)
	s.initCookieAuthenticator(s.cfg.Env)
	s.initRedisConnection(s.cfg.Redis)

	srv := &http.Server{
		Addr:    ":" + s.cfg.HTTPServer.Port,
		Handler: framework.Handler(s.cfg),
	}

	go func() {
		tracer.Start()

		if err := profiler.Start(); err != nil {
			log.WithError(err).Errorln("profiler fail")
		}

		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.WithError(err).Errorln("ListenAndServe fail")
			panic(err)
		}
	}()

	log.WithField("port", s.cfg.HTTPServer.Port).Debugln("server started")

	<-stopSignal

	tracer.Stop()
	profiler.Stop()

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Info("server stopped")

	return nil
}
