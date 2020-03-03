package redis

import (
	"errors"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

// Config config struct
type Config struct {
	URL         string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

// NewConfig new config
func NewConfig() (*Config, error) {
	// New config
	config := Config{
		URL: os.Getenv("REDIS_URL"),
	}
	// Validate driver
	if err := config.Validate(); err != nil {
		return nil, err
	}
	// Get config from env
	config.GetMaxIdleFromEnv()
	config.GetMaxActiveFromEnv()
	config.GetIdleTimeoutFromEnv()

	return &config, nil
}

// GetMaxIdleFromEnv get max idle from env
func (c *Config) GetMaxIdleFromEnv() {
	if maxIdle, err := strconv.Atoi(os.Getenv("REDIS_MAX_IDLE")); err == nil {
		c.MaxIdle = maxIdle
	}
}

// GetMaxActiveFromEnv get max idle from env
func (c *Config) GetMaxActiveFromEnv() {
	if maxActive, err := strconv.Atoi(os.Getenv("REDIS_MAX_ACTIVE")); err == nil {
		c.MaxActive = maxActive
	}
}

// GetIdleTimeoutFromEnv get max idle from env
func (c *Config) GetIdleTimeoutFromEnv() {
	if idleTimeout, err := strconv.Atoi(os.Getenv("REDIS_IDLE_TIMEOUT")); err == nil {
		c.IdleTimeout = idleTimeout
	}
}

// Validate validate
func (c *Config) Validate() error {
	if len(c.URL) == 0 {
		return errors.New("url can't be blank")
	}

	return nil
}

// LogrusFields logrus fields
func (c *Config) LogrusFields() logrus.Fields {
	return logrus.Fields{
		"url":          c.URL,
		"max_idle":     c.MaxIdle,
		"max_active":   c.MaxActive,
		"idle_timeout": c.IdleTimeout,
	}
}
