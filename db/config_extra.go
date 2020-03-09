package db

import (
	"os"
	"strconv"
)

// DefaultValues
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
