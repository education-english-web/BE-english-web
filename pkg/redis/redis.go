// Package redis abstracts Redis or a key-value database implementation
package redis

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis.v8"
)

var singleton *redis.Client

// Config redis for redis server
type Config struct {
	ConnectionString   string // if this field is non empty then other fields will be ignore when building connection string
	Host               string
	Port               int
	Database           int
	User               string
	Pass               string
	PoolSize           int
	IdleTimeout        int
	ReadTimeout        int
	WriteTimeout       int
	MinIdleConns       int
	InsecureSkipVerify bool
	TLS                bool
}

// ParseConfig return redis config from connection string
func (c *Config) ParseConfig() error {
	if c.ConnectionString != "" {
		connStr := strings.TrimPrefix(c.ConnectionString, "redis://")
		re := regexp.MustCompile(`(?P<user>[^\:]+)\:(?P<pass>[^@]*)@(?P<host>[^:]+):(?P<port>\d+)`)

		matches := re.FindStringSubmatch(connStr)
		userIndex := re.SubexpIndex("user")
		passIndex := re.SubexpIndex("pass")
		hostIndex := re.SubexpIndex("host")
		portIndex := re.SubexpIndex("port")

		if userIndex < 0 || passIndex < 0 || hostIndex < 0 || portIndex < 0 {
			return errors.New("invalid connection string")
		}

		c.Host = matches[hostIndex]
		c.Port, _ = strconv.Atoi(matches[portIndex])
		c.User = matches[userIndex]
		c.Pass = matches[passIndex]
	}

	return nil
}

// Setup creates redis client
func Setup(c Config) error {
	if singleton != nil {
		return nil
	}

	var tlsConfig *tls.Config

	if c.TLS {
		// #nosec
		//nolint:gosec
		tlsConfig = &tls.Config{
			InsecureSkipVerify: c.InsecureSkipVerify,
			VerifyConnection: func(cs tls.ConnectionState) error {
				opts := x509.VerifyOptions{
					DNSName:       cs.ServerName,
					Intermediates: x509.NewCertPool(),
				}

				for _, cert := range cs.PeerCertificates[1:] {
					opts.Intermediates.AddCert(cert)
				}

				_, err := cs.PeerCertificates[0].Verify(opts)

				return err
			},
		}
	}

	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", c.Host, c.Port), // use default Addr
		Password:     c.Pass,                               // no password set
		DB:           c.Database,                           // use default DB
		PoolSize:     c.PoolSize,
		ReadTimeout:  time.Duration(c.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(c.WriteTimeout) * time.Second,
		MinIdleConns: c.MinIdleConns,
		TLSConfig:    tlsConfig,
	})

	redistrace.WrapClient(client, redistrace.WithServiceName("stampless_redis"))

	if err := client.Ping(context.Background()).Err(); err != nil {
		return fmt.Errorf("error while pinging redis: %w", err)
	}

	singleton = client

	return nil
}

// Get gets the instance of singleton
func Get() *redis.Client {
	return singleton
}
