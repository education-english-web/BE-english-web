package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

// APP env: should only put environment variable related to the service itself, e.g.: Application Name, version, running environment, ...
var (
	EnvFlag = &cli.StringFlag{
		Name:    "env",
		Usage:   "Application environment: development, staging, production",
		EnvVars: []string{"DD_ENV"},
		Value:   os.Getenv("DD_ENV"),
	}

	AppNameFlag = &cli.StringFlag{
		Name:    "app_name",
		Usage:   "Application name",
		EnvVars: []string{"DD_SERVICE"},
		Value:   "Education_backend",
	}

	AppVersionFlag = &cli.StringFlag{
		Name:    "app_version",
		Usage:   "Application version",
		EnvVars: []string{"DD_VERSION"},
		Value:   "v1",
	}

	HTTPPortFlag = &cli.StringFlag{
		Name:    "http_port",
		Usage:   "Port binding to application",
		EnvVars: []string{"HTTP_PORT"},
		Value:   os.Getenv("HTTP_PORT"),
	}
	ProxyURLFlag = &cli.StringFlag{
		Name:    "proxy_url",
		Usage:   "Proxy URL",
		EnvVars: []string{"PROXY_URL"},
		Value:   "",
	}
)

// Postgres env
var (
	PostgresConnFlag = &cli.StringFlag{
		Name:    "postgres_conn",
		Usage:   `specify Postgres connection string. If non-empty then other flags begin with "postgres_" will be ignore`,
		EnvVars: []string{"DATABASE_URL"}, // support for heroku deployment.
		Value:   "",
	}

	PostgresHostFlag = &cli.StringFlag{
		Name:    "postgres_host",
		Usage:   "specify Postgres host",
		EnvVars: []string{"POSTGRES_HOST"},
	}

	PostgresPortFlag = &cli.StringFlag{
		Name:    "postgres_port",
		Usage:   "Postgres port is using by application",
		EnvVars: []string{"POSTGRES_PORT"},
	}

	PostgresMasterHostsFlag = &cli.StringFlag{
		Name:    "postgres_master_hosts",
		Usage:   "specify Postgres master hosts with port",
		EnvVars: []string{"POSTGRES_MASTER_HOSTS"},
	}

	PostgresSlaveHostsFlag = &cli.StringFlag{
		Name:    "postgres_slave_hosts",
		Usage:   "specify Postgres slave hosts with port",
		EnvVars: []string{"POSTGRES_SLAVE_HOSTS"},
		Value:   os.Getenv("POSTGRES_SLAVE_HOSTS"),
	}

	PostgresUserFlag = &cli.StringFlag{
		Name:    "postgres_user",
		Usage:   "specify Postgres user",
		EnvVars: []string{"POSTGRES_USER"},
	}

	PostgresPasswordFlag = &cli.StringFlag{
		Name:    "postgres_password",
		Usage:   "password used for Postgres user",
		EnvVars: []string{"POSTGRES_PASSWORD"},
	}

	PostgresDatabaseFlag = &cli.StringFlag{
		Name:    "postgres_db",
		Usage:   "Postgres database is using by application",
		EnvVars: []string{"POSTGRES_DB"},
	}

	PostgresMaxOpenConnsFlag = &cli.IntFlag{
		Name:    "postgres_max_open_conns",
		Usage:   "sets the maximum number of open connections to the database",
		EnvVars: []string{"POSTGRES_MAX_OPEN_CONNS"},
		Value:   50,
	}

	PostgresMaxIdleConnsFlag = &cli.IntFlag{
		Name:    "postgres_max_idle_conns",
		Usage:   "sets the maximum number of connections in the idle connection pool",
		EnvVars: []string{"POSTGRES_MAX_IDLE_CONNS"},
		Value:   5,
	}

	PostgresConnMaxLifetimeFlag = &cli.IntFlag{
		Name:    "postgres_conn_max_lifetime",
		Usage:   "sets the maximum amount of time in minutes a connection may be reused",
		EnvVars: []string{"POSTGRES_CONN_MAX_LIFETIME"},
		Value:   60,
	}
)

// For Redis
var (
	RedisConnFlag = &cli.StringFlag{
		Name:    "redis_conn",
		Usage:   `specify Redis connection string. If non empty then other flags begin with "redis_" will be ignore`,
		EnvVars: []string{"REDIS_CONN"}, // support for heroku deployment
		Value:   "",
	}

	RedisHostFlag = &cli.StringFlag{
		Name:    "redis_host",
		Usage:   "specify Redis host",
		EnvVars: []string{"REDIS_HOST"},
	}

	RedisPortFlag = &cli.StringFlag{
		Name:    "redis_port",
		Usage:   "Redis port is using by application",
		EnvVars: []string{"REDIS_PORT"},
	}

	RedisUserFlag = &cli.StringFlag{
		Name:    "redis_user",
		Usage:   "specify Redis user",
		EnvVars: []string{"REDIS_USER"},
		Value:   "default",
	}

	RedisPasswordFlag = &cli.StringFlag{
		Name:    "redis_password",
		Usage:   "password used for Redis user",
		EnvVars: []string{"REDIS_PASSWORD"},
	}

	RedisEnabledTLSFlag = &cli.BoolFlag{
		Name:    "redis_enabled_tls",
		Usage:   "enable tls for Redis tls connection",
		EnvVars: []string{"REDIS_ENABLED_TLS"},
		Value:   false,
	}

	RedisInsecureSkipVerifyFlag = &cli.BoolFlag{
		Name:    "redis_insecure_skip_verify",
		Usage:   "insecure_skip_verify used for Redis tls verify",
		EnvVars: []string{"REDIS_INSECURE_SKIP_VERIFY"},
		Value:   true,
	}

	RedisDatabaseFlag = &cli.IntFlag{
		Name:    "redis_db",
		Usage:   "Redis database is using by application",
		EnvVars: []string{"REDIS_DB"},
		Value:   0,
	}

	RedisPoolSizeFlag = &cli.IntFlag{
		Name:    "redis_max_open_conns",
		Usage:   "sets the maximum number of open connections to the database",
		EnvVars: []string{"REDIS_POOL_SIZE"},
		Value:   10,
	}
)

// For Pusher
var (
	PusherAppIDFlag = &cli.StringFlag{
		Name:    "pusher_app_id",
		Usage:   "pusher app_id",
		EnvVars: []string{"PUSHER_APP_ID"},
		Value:   "1597954",
	}

	PusherKeyFlag = &cli.StringFlag{
		Name:    "pusher_key",
		Usage:   "pusher key",
		EnvVars: []string{"PUSHER_KEY"},
		Value:   "54461a4353c6e24e361d",
	}

	PusherSecretFlag = &cli.StringFlag{
		Name:    "pusher_secret",
		Usage:   "pusher secret",
		EnvVars: []string{"PUSHER_SECRET"},
		Value:   "7c218d7442e0da6eee5c",
	}

	PusherClusterFlag = &cli.StringFlag{
		Name:    "pusher_cluster",
		Usage:   "pusher cluster",
		EnvVars: []string{"PUSHER_CLUSTER"},
		Value:   "ap1",
	}

	PusherSecureFlag = &cli.BoolFlag{
		Name:    "pusher_secure",
		Usage:   "pusher secure",
		EnvVars: []string{"PUSHER_SECURE"},
		Value:   true,
	}

	PusherEncryptionMasterKeyBase64Flag = &cli.StringFlag{
		Name:    "pusher_encryption_master_key_base64",
		Usage:   "pusher encryption master key base64",
		EnvVars: []string{"PUSHER_ENCRYPTION_MASTER_KEY_BASE64"},
		Value:   "6Jlfv99Ex7rebcWdn7U7a77Gqb+/3ZowXrutTFo0k8w=",
	}
)

// LocalDataDirFlag Local storage
var (
	LocalDataDirFlag = &cli.StringFlag{
		Name:    "local_data_dir",
		Usage:   "local data directory path",
		EnvVars: []string{"LOCAL_DATA_DIR"},
		Value:   "assets",
	}
)

// Log and notifier env
var (
	LogLevelFlag = &cli.StringFlag{
		Name:    "log_level",
		Usage:   "Level to log message to standard logger: panic, fatal, error, warn, warning, info, debug",
		EnvVars: []string{"LOG_LEVEL"},
		Value:   "debug",
	}

	RollbarTokenFlag = &cli.StringFlag{
		Name:    "rollbar_token",
		Usage:   "Token to access rollbar service",
		EnvVars: []string{"ROLLBAR_TOKEN"},
		Value:   "",
	}

	NotifierEngineFlag = &cli.StringFlag{
		Name:    "notifier_engine",
		Usage:   "Define notifier engine to use: rollbar",
		EnvVars: []string{"NOTIFIER_ENGINE"},
		Value:   "",
	}
)

// Monitoring tool env
var (
	DatadogAgentHostFlag = &cli.StringFlag{
		Name:    "dd_agent_host",
		Usage:   "Define Datadog agent host",
		EnvVars: []string{"DD_AGENT_HOST"},
		Value:   "localhost",
	}

	DatadogAgentAPMPortFlag = &cli.StringFlag{
		Name:    "dd_agent_apm_port",
		Usage:   "Define Datadog agent port",
		EnvVars: []string{"DD_AGENT_APM_PORT"},
		Value:   "8126",
	}

	TracerEngineFlag = &cli.StringFlag{
		Name:    "tracer_engine",
		Usage:   "Define logger engine to use: datadog",
		EnvVars: []string{"TRACER_ENGINE"},
		Value:   "",
	}

	ProfilerEngineFlag = &cli.StringFlag{
		Name:    "profiler_engine",
		Usage:   "Define profiler engine to use: datadog",
		EnvVars: []string{"PROFILER_ENGINE"},
		Value:   "",
	}
)

// EnabledProfilingFlag For pprof middleware
var (
	EnabledProfilingFlag = &cli.BoolFlag{
		Name:    "enabled_pprof",
		Usage:   "enable pprof middleware",
		EnvVars: []string{"ENABLED_PPROF"},
		Value:   false,
	}
)

// CORSAllowHostsFlag CORS env
var (
	CORSAllowHostsFlag = &cli.StringFlag{
		Name:    "cors_allow_hosts",
		Usage:   "List of origins a cross-domain request can be executed from, separated by comma",
		EnvVars: []string{"CORS_ALLOW_HOSTS"},
		Value:   "http://localhost:30003,https://stampless-frontend.herokuapp.com,http://127.0.0.1:5500",
	}
)

// For related Aweb configuration
var (
	GSuiteClientIDFlag = &cli.StringFlag{
		Name:    "gsuite_client_id",
		Usage:   "client id of gsuite oauth application",
		EnvVars: []string{"GSUITE_CLIENT_ID"},
		Value:   "",
	}

	GSuiteClientSecretFlag = &cli.StringFlag{
		Name:    "gsuite_client_secret",
		Usage:   "client secret of gsuite oauth application",
		EnvVars: []string{"GSUITE_CLIENT_SECRET"},
		Value:   "",
	}

	OIDCAllowDomains = &cli.StringFlag{
		Name:    "oidc_allow_domains",
		Usage:   "Domains are allowed when using oidc",
		EnvVars: []string{"OIDC_ALLOW_DOMAINS"},
		Value:   "moneyforward.vn,moneyforward.co.jp",
	}
)

// Google service account related env vars
var (
	GServiceAccountEmailFlag = &cli.StringFlag{
		Name:     "gservice_account_email",
		Usage:    "google service account email",
		EnvVars:  []string{"GSERVICE_ACCOUNT_EMAIL"},
		Value:    "",
		Required: true,
	}

	GServiceAccountPrivateKeyFlag = &cli.StringFlag{
		Name:     "gservice_account_private_key",
		Usage:    "google service account private key",
		EnvVars:  []string{"GSERVICE_ACCOUNT_PRIVATE_KEY"},
		Value:    "",
		Required: true,
	}

	GServiceAccountPrivateKeyIDFlag = &cli.StringFlag{
		Name:     "gservice_account_private_key_id",
		Usage:    "google servie account private key id",
		EnvVars:  []string{"GSERVICE_ACCOUNT_PRIVATE_KEY_ID"},
		Value:    "",
		Required: true,
	}
	GServiceSpreadSheetIDOfficesReportFlag = &cli.StringFlag{
		Name:     "gservice_spreadsheet_id_offices_report",
		Usage:    "offices report google spreadsheet id",
		EnvVars:  []string{"GSERVICE_SPREADSHEET_ID_OFFICES_REPORT"},
		Value:    "",
		Required: true,
	}
	GServiceSpreadSheetIDOfficesStatusReportFlag = &cli.StringFlag{
		Name:     "gservice_spreadsheet_id_offices_status_report",
		Usage:    "offices status report google spreadsheet id",
		EnvVars:  []string{"GSERVICE_SPREADSHEET_ID_OFFICES_STATUS_REPORT"},
		Value:    "",
		Required: true,
	}
)

// Salt for password
var (
	SaltFlag = &cli.StringFlag{
		Name:    "salt",
		Usage:   "salt for password",
		EnvVars: []string{"SALT"},
	}
)

// JWTSecretFlag for jwt
var (
	JWTSecretFlag = &cli.StringFlag{
		Name:    "jwt_secret",
		Usage:   "jwt",
		EnvVars: []string{"JWT_SECRET"},
	}
)
