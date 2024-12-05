package config

import (
	"fmt"
	"strings"

	"github.com/education-english-web/BE-english-web/pkg/redis"
)

const (
	ENVProduction  = "production"
	ENVStaging     = "staging"
	ENVHeroku      = "heroku"
	ENVDevelopment = "development"
)

// CORS hold configuration of Cross-origin resource sharing
type CORS struct {
	AllowHosts []string
}

// Postgres config for Postgres server
type Postgres struct {
	ConnectionString string // if this field is non-empty then other fields will be ignored when building connection string
	Masters          string
	Slaves           string
	User             string
	Password         string
	DB               string
	Port             int
	MaxOpenConns     int
	MaxIdleConns     int
	ConnMaxLifetime  int // time in minute
	IsEnabledLog     bool
}

// LocalStorage config for local file store
type LocalStorage struct {
	DataDir string
}

// ServerAddr server addresses
type ServerAddr struct {
	Port string
}

// Rollbar hold configuration of Rollbar reporting
type Rollbar struct {
	Token string
	Env   string
}

// Datadog hold configuration of Datadog tracer
type Datadog struct {
	Env            string
	ServiceName    string
	ServiceVersion string
	Host           string
	AgentAMPPort   string
}

// GSuite client information
type GSuite struct {
	ClientID     string
	ClientSecret string
}

// GServiceAccount Google service account config
type GServiceAccount struct {
	Email                            string
	PrivateKey                       string
	PrivateKeyID                     string
	SpreadSheetIDOfficesReport       string
	SpreadSheetIDOfficesStatusReport string
}

// Salesforce related env variables
type Salesforce struct {
	ClientID     string
	ClientSecret string
	PrivateKey   string
}

// Config is APP config information
type Config struct {
	Env              string
	HTTPServer       ServerAddr
	CORS             CORS
	Postgres         Postgres
	Redis            redis.Config
	LocalStorage     LocalStorage
	Rollbar          Rollbar
	NotifierEngine   string
	LogLevel         string
	Datadog          Datadog
	TracerEngine     string
	ProfilerEngine   string
	JWTSecret        string
	EnabledProfiling bool
	FQDN             string
	ProxyURL         string
	GSuite           GSuite
	OIDCAllowDomains []string
	Salesforce       Salesforce
	GServiceAccount  GServiceAccount
	Salt             string
}

// Conn return connection string
func (p *Postgres) Conn(host string) string {
	if p.ConnectionString != "" {
		// Trim prefix and split the connection string from query parameters
		connStr := strings.TrimPrefix(p.ConnectionString, "postgres://")
		connStr = strings.Split(connStr, "?")[0]
		// Replace '@' with '@' (to avoid issues) and ':' for port if not present
		connStr = strings.Replace(connStr, "@", "@", 1)
		connStr = strings.Replace(connStr, ":", ":", 1)
		// Ensure that port is correctly set and database name is included
		if !strings.Contains(connStr, ":5432") {
			connStr = strings.Replace(connStr, "/", ":5432/", 1)
		}
		connStr = strings.Replace(connStr, "/", "/?sslmode=disable", 1)
		return connStr
	}

	// Correct DSN format for PostgreSQL
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Tokyo",
		host, p.Port, p.User, p.Password, p.DB)
}
