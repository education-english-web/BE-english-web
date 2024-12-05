package main

import (
	"os"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"

	"github.com/education-english-web/BE-english-web/app/config"
	//"github.com/education-english-web/BE-english-web/pkg/dynamodbmanager"
	"github.com/education-english-web/BE-english-web/pkg/gormutil"
	"github.com/education-english-web/BE-english-web/pkg/log"
	"github.com/education-english-web/BE-english-web/pkg/redis"
)

const ApplicationLoadFail = 1

func (s *service) initDBConnectionPostgres(cfg config.Postgres) *gorm.DB {
	if cfg.Masters == "" {
		log.Errorln("miss db info")
		os.Exit(ApplicationLoadFail)
	}

	var (
		db         *gorm.DB
		gormConfig = gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
		cfgMasters       = strings.Split(cfg.Masters, ",")
		cfgSlaves        = strings.Split(cfg.Slaves, ",")
		masterDialectors = make([]gorm.Dialector, 0, len(cfgMasters))
		slaveDialectors  = make([]gorm.Dialector, 0, len(cfgSlaves))
	)

	if cfg.IsEnabledLog {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	master, otherMasters := cfgMasters[0], cfgMasters[1:]

	for _, otherMaster := range otherMasters {
		if otherMaster == "" {
			continue
		}

		dialector := postgres.New(postgres.Config{DSN: cfg.Conn(otherMaster)})
		masterDialectors = append(masterDialectors, dialector)
	}

	for _, slave := range cfgSlaves {
		if slave == "" {
			continue
		}

		dialector := postgres.New(postgres.Config{DSN: cfg.Conn(slave)})
		slaveDialectors = append(slaveDialectors, dialector)
	}

	log.Infof("Connecting to DB with DSN: %s", master)

	db, err := gormutil.OpenDBConnectionPostgreSQL(cfg.Conn(master), gormConfig)
	if err != nil {
		log.WithError(err).
			Errorln("creating connection to DB")
		os.Exit(ApplicationLoadFail)
	}

	resolver := dbresolver.Register(dbresolver.Config{
		Sources:           masterDialectors,
		Replicas:          slaveDialectors,
		Policy:            dbresolver.RandomPolicy{},
		TraceResolverMode: cfg.IsEnabledLog,
	})

	resolver.SetMaxOpenConns(cfg.MaxOpenConns)
	resolver.SetMaxIdleConns(cfg.MaxIdleConns)
	resolver.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Minute)

	if err := db.Use(resolver); err != nil {
		log.WithError(err).
			WithField("master_dialectors", masterDialectors).
			WithField("slave_dialectors", slaveDialectors).
			Errorln("fail to register master slave dbs")
		os.Exit(ApplicationLoadFail)
	}

	rawDB, err := db.DB()
	if err != nil {
		log.WithError(err).Errorln("get DB failed")
		os.Exit(ApplicationLoadFail)
	}

	rawDB.SetMaxOpenConns(cfg.MaxOpenConns)
	rawDB.SetMaxIdleConns(cfg.MaxIdleConns)
	rawDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Minute)

	return db
}

func (s *service) initRedisConnection(cfg redis.Config) {
	if err := cfg.ParseConfig(); err != nil {
		log.WithError(err).Errorln("fail to parse redis configuration")
		os.Exit(ApplicationLoadFail)
	}

	if err := redis.Setup(cfg); err != nil {
		log.WithError(err).Errorln("fail to connect redis")
		os.Exit(ApplicationLoadFail)
	}
}

// func (s *service) initDynamoDBConnection(cfg config.AWS) {
// 	if err := dynamodbmanager.Setup(cfg.Region, cfg.Endpoint); err != nil {
// 		log.WithError(err).Errorln("fail to connect dynamo db")
// 		os.Exit(ApplicationLoadFail)
// 	}
// }
