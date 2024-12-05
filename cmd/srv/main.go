package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	flags := []cli.Flag{
		// Common
		EnvFlag,
		AppNameFlag,
		AppVersionFlag,

		// Postgres
		PostgresConnFlag,
		PostgresMasterHostsFlag,
		PostgresSlaveHostsFlag,
		PostgresUserFlag,
		PostgresHostFlag,
		PostgresPortFlag,
		PostgresPasswordFlag,
		PostgresDatabaseFlag,
		PostgresMaxOpenConnsFlag,
		PostgresMaxIdleConnsFlag,
		PostgresConnMaxLifetimeFlag,

		// Redis
		RedisConnFlag,
		RedisHostFlag,
		RedisPortFlag,
		RedisUserFlag,
		RedisPasswordFlag,
		RedisDatabaseFlag,
		RedisPoolSizeFlag,
		RedisInsecureSkipVerifyFlag,
		RedisEnabledTLSFlag,

		PusherAppIDFlag,
		PusherKeyFlag,
		PusherSecretFlag,
		PusherClusterFlag,
		PusherSecureFlag,
		PusherEncryptionMasterKeyBase64Flag,

		LocalDataDirFlag,
		HTTPPortFlag,
		RollbarTokenFlag,
		NotifierEngineFlag,
		LogLevelFlag,
		DatadogAgentHostFlag,
		DatadogAgentAPMPortFlag,
		TracerEngineFlag,
		ProfilerEngineFlag,
		EnabledProfilingFlag,
		CORSAllowHostsFlag,

		SaltFlag,
		JWTSecretFlag,
	}

	app := &cli.App{
		Name:  "Education Service",
		Flags: flags,
		Action: func(ctx *cli.Context) error {
			srv := newService(ctx)

			return srv.start()
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
