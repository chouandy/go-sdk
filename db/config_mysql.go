package db

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"

	"github.com/sirupsen/logrus"
)

// MySQLConfig mysql config struct
type MySQLConfig struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Charset  string
	Collate  string
}

// NewMySQLConfigFromEnvs new mysql config from envs
func NewMySQLConfigFromEnvs() (Config, error) {
	// New configs
	config := &MySQLConfig{
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Charset:  os.Getenv("DB_CHARSET"),
		Collate:  os.Getenv("DB_COLLATE"),
	}
	// Validate driver
	if err := config.Validate(); err != nil {
		return nil, err
	}
	// Load default
	config.LoadDefault()

	return config, nil
}

// NewMySQLConfigFromDatabaseURL new mysql config from database url
func NewMySQLConfigFromDatabaseURL() (Config, error) {
	// Get database url from env
	databaseURL := os.Getenv("DATABASE_URL")

	// Parse database url
	cfg, err := mysql.ParseDSN(strings.TrimPrefix(databaseURL, "mysql://"))
	if err != nil {
		return nil, err
	}

	// New configs
	config := &MySQLConfig{
		Driver:   "mysql",
		Database: cfg.DBName,
		Username: cfg.User,
		Password: cfg.Passwd,
		Collate:  cfg.Collation,
	}
	tmp := strings.Split(cfg.Addr, ":")
	config.Host = tmp[0]
	if len(tmp) > 1 {
		config.Port = tmp[1]
	}

	// Validate
	if err := config.Validate(); err != nil {
		return nil, err
	}

	// Load default
	config.LoadDefault()

	return config, nil
}

// Validate validate
func (c *MySQLConfig) Validate() error {
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
func (c *MySQLConfig) LoadDefault() {
	// Set port default value
	if len(c.Port) == 0 {
		c.Port = "3306"
	}
	// Set charset default value
	if len(c.Charset) == 0 {
		c.Charset = "utf8"
	}
	// Set default collate default value
	if len(c.Collate) == 0 {
		c.Collate = "utf8mb4_general_ci"
	}
}

// GetDriver get driver
func (c *MySQLConfig) GetDriver() string {
	return c.Driver
}

// DatabaseURL database url
func (c *MySQLConfig) DatabaseURL() string {
	// New query
	query := url.Values{}
	query.Set("charset", c.Charset)
	query.Set("parseTime", "true")

	return fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?%s",
		c.Username, c.Password, c.Host, c.Port, c.Database, query.Encode(),
	)
}

// DataSource data source
func (c *MySQLConfig) DataSource() string {
	// New query
	query := url.Values{}
	query.Set("charset", c.Charset)
	query.Set("parseTime", "true")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		c.Username, c.Password, c.Host, c.Port, c.Database, query.Encode(),
	)
}

// DataSourceWithoutDatabase data source without database
func (c *MySQLConfig) DataSourceWithoutDatabase() string {
	// New query
	query := url.Values{}
	query.Set("charset", c.Charset)

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/?%s",
		c.Username, c.Password, c.Host, c.Port, query.Encode(),
	)
}

// CreateDatabaseStatement create database statement
func (c *MySQLConfig) CreateDatabaseStatement() string {
	return fmt.Sprintf("CREATE DATABASE `%s` DEFAULT CHARACTER SET = '%s' DEFAULT COLLATE '%s';",
		c.Database, c.Charset, c.Collate,
	)
}

// DropDatabaseStatement drop database statement
func (c *MySQLConfig) DropDatabaseStatement() string {
	return fmt.Sprintf("DROP DATABASE `%s`", c.Database)
}

// LogrusFields logrus fields
func (c *MySQLConfig) LogrusFields() logrus.Fields {
	return logrus.Fields{
		"driver":   c.Driver,
		"host":     c.Host,
		"port":     c.Port,
		"database": c.Database,
	}
}
