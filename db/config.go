package db

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

// DataSourceFormat data source format
var DataSourceFormat = map[string]string{
	"mysql":    "%s:%s@tcp(%s:%s)/%s?%s",
	"postgres": "postgres://%s:%s@%s:%s/%s?%s",
}

// DataSourceWithoutDatabaseFormat data source without database format
var DataSourceWithoutDatabaseFormat = map[string]string{
	"mysql":    "%s:%s@tcp(%s:%s)/?%s",
	"postgres": "postgres://%s:%s@%s:%s/?%s",
}

// CreateDatabaseStatementFormat create database statement format
var CreateDatabaseStatementFormat = map[string]string{
	"mysql":    "CREATE DATABASE `%s` DEFAULT CHARACTER SET = '%s' DEFAULT COLLATE '%s';",
	"postgres": `CREATE DATABASE "%s" ENCODING = '%s' LC_COLLATE = '%s' LC_CTYPE = '%s';`,
}

// DropDatabaseStatementFormat drop database statement
var DropDatabaseStatementFormat = map[string]string{
	"mysql":    "DROP DATABASE `%s`",
	"postgres": `DROP DATABASE "%s"`,
}

// DatabaseURLFormat data source format
var DatabaseURLFormat = map[string]string{
	"mysql":    "mysql://%s:%s@tcp(%s:%s)/%s?%s",
	"postgres": "postgres://%s:%s@%s:%s/%s?%s",
}

// Config config struct
type Config struct {
	Driver         string
	Host           string
	Port           string
	Username       string
	Password       string
	Database       string
	Charset        string
	DefaultCollate string
	MaxOpenConns   int
	MaxIdleConns   int
	LogMode        bool
	SSLMode        string
}

// NewConfig new config
func NewConfig() (*Config, error) {
	// New config
	config := Config{
		Driver:         os.Getenv("DB_DRIVER"),
		Host:           os.Getenv("DB_HOST"),
		Port:           os.Getenv("DB_PORT"),
		Database:       os.Getenv("DB_DATABASE"),
		Username:       os.Getenv("DB_USERNAME"),
		Password:       os.Getenv("DB_PASSWORD"),
		Charset:        os.Getenv("DB_CHARSET"),
		DefaultCollate: os.Getenv("DB_DEFAULT_COLLATE"),
		SSLMode:        os.Getenv("DB_SSL_MODE"),
	}
	// Validate driver
	if err := config.Validate(); err != nil {
		return nil, err
	}
	// Get max open conns from env
	config.GetMaxOpenConnsFromEnv()
	// Get max idle conns from env
	config.GetMaxIdleConnsFromEnv()
	// Get log mode from env
	config.GetLogModeFromEnv()
	// Load default
	config.LoadDefault()

	return &config, nil
}

// Validate validate
func (c *Config) Validate() error {
	if len(c.Driver) == 0 {
		return errors.New("driver can't be blank")
	}
	if len(c.Host) == 0 {
		return errors.New("host can't be blank")
	}
	if len(c.Database) == 0 {
		return errors.New("database can't be blank")
	}

	return nil
}

// GetMaxOpenConnsFromEnv get max open conns from env
func (c *Config) GetMaxOpenConnsFromEnv() {
	if maxOpenConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS")); err == nil {
		c.MaxOpenConns = maxOpenConns
	}
}

// GetMaxIdleConnsFromEnv get max idle conns from env
func (c *Config) GetMaxIdleConnsFromEnv() {
	if maxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS")); err == nil {
		c.MaxIdleConns = maxIdleConns
	}
}

// GetLogModeFromEnv get log mode from env
func (c *Config) GetLogModeFromEnv() {
	if logMode, err := strconv.ParseBool(os.Getenv("DB_LOG_MODE")); err == nil {
		c.LogMode = logMode
	}
}

// LoadDefault load default
func (c *Config) LoadDefault() {
	// Postgres default values
	if c.Driver == "postgres" {
		// Set port default value
		if len(c.Port) == 0 {
			c.Port = "5432"
		}
		// Set charset default value
		if len(c.Charset) == 0 {
			c.Charset = "UTF8"
		}
		// Set default collate default value
		if len(c.DefaultCollate) == 0 {
			c.DefaultCollate = "en_US.utf8"
		}
		// Set ssl mode default value
		if c.SSLMode == "" {
			c.SSLMode = "disable"
		}
	} else {
		// Set port default value
		if len(c.Port) == 0 {
			c.Port = "3306"
		}
		// Set charset default value
		if len(c.Charset) == 0 {
			c.Charset = "utf8"
		}
		// Set default collate default value
		if len(c.DefaultCollate) == 0 {
			c.DefaultCollate = "utf8_general_ci"
		}
	}

	// Set max idle conns default value
	if c.MaxIdleConns == 0 {
		c.MaxIdleConns = 30
	}
	// Set max open conns default value
	if c.MaxOpenConns == 0 {
		c.MaxOpenConns = 150
	}
}

// DataSource data source
func (c *Config) DataSource() string {
	// New query
	query := url.Values{}
	if c.Driver == "postgres" {
		query.Set("sslmode", c.SSLMode)
	} else {
		query.Set("charset", c.Charset)
		query.Set("parseTime", "true")
	}

	return fmt.Sprintf(DataSourceFormat[c.Driver],
		c.Username, c.Password, c.Host, c.Port, c.Database, query.Encode(),
	)
}

// DataSourceWithoutDatabase data source without database
func (c *Config) DataSourceWithoutDatabase() string {
	// New query
	query := url.Values{}
	if c.Driver == "postgres" {
		query.Set("sslmode", c.SSLMode)
	} else {
		query.Set("charset", c.Charset)
	}

	return fmt.Sprintf(DataSourceWithoutDatabaseFormat[c.Driver],
		c.Username, c.Password, c.Host, c.Port, query.Encode(),
	)
}

// CreateDatabaseStatement create database statement
func (c *Config) CreateDatabaseStatement() string {
	// New args
	var args []interface{}
	if c.Driver == "postgres" {
		args = []interface{}{c.Database, c.Charset, c.DefaultCollate, c.DefaultCollate}
	} else {
		args = []interface{}{c.Database, c.Charset, c.DefaultCollate}
	}

	return fmt.Sprintf(CreateDatabaseStatementFormat[c.Driver], args...)
}

// DropDatabaseStatement drop database statement
func (c *Config) DropDatabaseStatement() string {
	return fmt.Sprintf(DropDatabaseStatementFormat[c.Driver], c.Database)
}

// DatabaseURL database url
func (c *Config) DatabaseURL() string {
	// New query
	query := url.Values{}
	if c.Driver == "postgres" {
		query.Set("sslmode", c.SSLMode)
	} else {
		query.Set("charset", c.Charset)
		query.Set("parseTime", "true")
	}

	return fmt.Sprintf(DatabaseURLFormat[c.Driver],
		c.Username, c.Password, c.Host, c.Port, c.Database, query.Encode(),
	)
}

// LogrusFields logrus fields
func (c *Config) LogrusFields() logrus.Fields {
	return logrus.Fields{
		"driver":   c.Driver,
		"host":     c.Host,
		"port":     c.Port,
		"database": c.Database,
	}
}
