package db

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"

	"github.com/sirupsen/logrus"
)

// PostgresConfig postgres config struct
type PostgresConfig struct {
	Driver       string
	Host         string
	Port         string
	Username     string
	Password     string
	Database     string
	Charset      string
	MaxOpenConns int
	MaxIdleConns int
	LogMode      bool
	SSLMode      string
}

// NewPostgresConfigFromEnvs new postgres config from envs
func NewPostgresConfigFromEnvs() (Config, error) {
	// New config
	config := PostgresConfig{
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Charset:  os.Getenv("DB_CHARSET"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}
	// Validate driver
	if err := config.Validate(); err != nil {
		return nil, err
	}
	// Load default
	config.LoadDefault()

	return &config, nil
}

// NewPostgresConfigFromDatabaseURL new postgres config from database url
func NewPostgresConfigFromDatabaseURL() (Config, error) {
	// Get database url from env
	databaseURL := os.Getenv("DATABASE_URL")

	// Parse database url
	u, err := url.Parse(databaseURL)
	if err != nil {
		return nil, err
	}

	if u.Scheme != "postgres" && u.Scheme != "postgresql" {
		return nil, fmt.Errorf("invalid connection protocol: %s", u.Scheme)
	}

	// New config
	config := PostgresConfig{
		Driver: "postgres",
	}

	if u.User != nil {
		config.Username = u.User.Username()

		if password, ok := u.User.Password(); ok {
			config.Password = password
		}
	}

	if host, port, err := net.SplitHostPort(u.Host); err != nil {
		config.Host = u.Host
	} else {
		config.Host = host
		config.Port = port
	}

	if u.Path != "" {
		config.Database = u.Path[1:]
	}

	if sslmode := u.Query().Get("sslmode"); len(sslmode) > 0 {
		config.SSLMode = sslmode
	}

	// Validate
	if err := config.Validate(); err != nil {
		return nil, err
	}

	// Load default
	config.LoadDefault()

	return &config, nil
}

// Validate validate
func (c *PostgresConfig) Validate() error {
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

// LoadDefault load default
func (c *PostgresConfig) LoadDefault() {
	// Set port default value
	if len(c.Port) == 0 {
		c.Port = "5432"
	}
	// Set charset default value
	if len(c.Charset) == 0 {
		c.Charset = "UTF8"
	}
	// Set ssl mode default value
	if c.SSLMode == "" {
		c.SSLMode = "require"
	}
}

// GetDriver get driver
func (c *PostgresConfig) GetDriver() string {
	return c.Driver
}

// DatabaseURL database url
func (c *PostgresConfig) DatabaseURL() string {
	// New query
	query := url.Values{}
	query.Set("sslmode", c.SSLMode)

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s",
		c.Username, c.Password, c.Host, c.Port, c.Database, query.Encode(),
	)
}

// DataSource data source
func (c *PostgresConfig) DataSource() string {
	// New query
	query := url.Values{}
	query.Set("sslmode", c.SSLMode)

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s",
		c.Username, c.Password, c.Host, c.Port, c.Database, query.Encode(),
	)
}

// DataSourceWithoutDatabase data source without database
func (c *PostgresConfig) DataSourceWithoutDatabase() string {
	// New query
	query := url.Values{}
	query.Set("sslmode", c.SSLMode)

	return fmt.Sprintf("postgres://%s:%s@%s:%s/?%s",
		c.Username, c.Password, c.Host, c.Port, query.Encode(),
	)
}

// CreateDatabaseStatement create database statement
func (c *PostgresConfig) CreateDatabaseStatement() string {
	return fmt.Sprintf(`CREATE DATABASE "%s" ENCODING = '%s';`, c.Database, c.Charset)
}

// DropDatabaseStatement drop database statement
func (c *PostgresConfig) DropDatabaseStatement() string {
	return fmt.Sprintf(`DROP DATABASE "%s"`, c.Database)
}

// LogrusFields logrus fields
func (c *PostgresConfig) LogrusFields() logrus.Fields {
	return logrus.Fields{
		"driver":   c.Driver,
		"host":     c.Host,
		"port":     c.Port,
		"database": c.Database,
	}
}
