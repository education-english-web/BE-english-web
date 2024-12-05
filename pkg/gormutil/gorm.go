package gormutil

import (
	"fmt"
	"sync"

	"github.com/lib/pq"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	gormtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorm.io/gorm.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var (
	once      sync.Once
	singleton *gorm.DB
)

// OpenDBConnection opens a DB connection for PostgreSQL.
func OpenDBConnectionPostgreSQL(conn string, config gorm.Config) (*gorm.DB, error) {
	var err error

	once.Do(func() {
		sqltrace.Register("postgres", pq.Driver{}, sqltrace.WithServiceName("pakka-sql"))
		sqlDB, innerErr := sqltrace.Open("postgres", conn)
		if innerErr != nil {
			err = fmt.Errorf("sqltrace open: %w", innerErr)
			return
		}

		db, innerErr := gormtrace.Open(
			postgres.New(postgres.Config{Conn: sqlDB}),
			&config,
			gormtrace.WithServiceName("stampless-gorm"),
		)
		if innerErr != nil {
			err = fmt.Errorf("gormtrace open: %w", innerErr)
			return
		}

		singleton = db
	})

	if err != nil {
		return nil, err
	}

	return singleton, nil
}

// GetDB gets the instance of singleton
func GetDB() *gorm.DB {
	if singleton != nil {
		return singleton.
			Clauses(dbresolver.Write).
			Session(&gorm.Session{NewDB: true})
	}

	return nil
}
