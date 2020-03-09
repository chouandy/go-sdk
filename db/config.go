package db

import (
	"errors"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var config Config

var errUnknownDBDriver = errors.New("unknown db driver")

// Config config interface
type Config interface {
	GetDriver() string
	Validate() error
	LoadDefault()
	DatabaseURL() string
	DataSource() string
	DataSourceWithoutDatabase() string
	CreateDatabaseStatement() string
	DropDatabaseStatement() string
	LogrusFields() logrus.Fields
}

// NewConfig new config
func NewConfig() error {
	// New config from envs
	driver := os.Getenv("DB_DRIVER")
	if len(driver) > 0 {
		return NewConfigFromEnvs(driver)
	}

	// New config from database url
	databaseURL := os.Getenv("DATABASE_URL")
	if len(databaseURL) > 0 {
		return NewConfigFromDatabaseURL(databaseURL)
	}

	return errUnknownDBDriver
}

// NewConfigFromEnvs new config from envs
func NewConfigFromEnvs(driver string) (err error) {
	if driver == "mysql" {
		config, err = NewMySQLConfigFromEnvs()
		return
	} else if driver == "postgres" {
		config, err = NewPostgresConfigFromEnvs()
		return
	}

	return errUnknownDBDriver
}

// NewConfigFromDatabaseURL new config from database url
func NewConfigFromDatabaseURL(databaseURL string) (err error) {
	if strings.HasPrefix(databaseURL, "mysql://") {
		config, err = NewMySQLConfigFromDatabaseURL()
		return
	} else if strings.HasPrefix(databaseURL, "postgres://") {
		config, err = NewPostgresConfigFromDatabaseURL()
		return
	}

	return errUnknownDBDriver
}

// GetConfig get config
func GetConfig() Config {
	return config
}
