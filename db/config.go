package db

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

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
func NewConfig() (Config, error) {
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

	return nil, errUnknownDBDriver
}

// NewConfigFromEnvs new config from envs
func NewConfigFromEnvs(driver string) (Config, error) {
	if driver == "mysql" {
		return NewMySQLConfigFromEnvs()
	} else if driver == "postgres" {
		return NewPostgresConfigFromEnvs()
	}

	return nil, errUnknownDBDriver
}

// NewConfigFromDatabaseURL new config from database url
func NewConfigFromDatabaseURL(databaseURL string) (Config, error) {
	if strings.HasPrefix(databaseURL, "mysql://") {
		return NewMySQLConfigFromDatabaseURL()
	} else if strings.HasPrefix(databaseURL, "postgres://") {
		return NewPostgresConfigFromDatabaseURL()
	}

	return nil, errUnknownDBDriver
}

// DefaultValue
const (
	DefaultMaxIdleConns = 30
	DefaultMaxOpenConns = 150
)

// GetMaxOpenConnsFromEnv get max open conns from env
func GetMaxOpenConnsFromEnv() int {
	maxOpenConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	if err != nil {
		maxOpenConns = DefaultMaxOpenConns
	}

	return maxOpenConns
}

// GetMaxIdleConnsFromEnv get max idle conns from env
func GetMaxIdleConnsFromEnv() int {
	maxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	if err != nil {
		maxIdleConns = DefaultMaxIdleConns
	}

	return maxIdleConns
}

// GetLogModeFromEnv get log mode from env
func GetLogModeFromEnv() bool {
	logMode, _ := strconv.ParseBool(os.Getenv("DB_LOG_MODE"))
	return logMode
}
