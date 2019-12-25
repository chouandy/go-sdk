package redis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"

	dgolog "github.com/chouandy/go-sdk/log"
)

var pool *redis.Pool

// Config config struct
type Config struct {
	Host        string `json:"host" yaml:"host"`
	Port        int    `json:"port" yaml:"port"`
	MaxIdle     int    `json:"max_idle" yaml:"max_idle"`
	MaxActive   int    `json:"max_active" yaml:"max_active"`
	IdleTimeout int    `json:"idle_timeout" yaml:"idle_timeout"`
}

// LogrusFields logrus fields
func (c *Config) LogrusFields() logrus.Fields {
	return logrus.Fields{
		"host":         c.Host,
		"port":         c.Port,
		"max_idle":     c.MaxIdle,
		"max_active":   c.MaxActive,
		"idle_timeout": c.IdleTimeout,
	}
}

// Address address
func (c *Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// Init init redis pool
func Init(config Config) {
	// Print log
	dgolog.TextLog().WithFields(config.LogrusFields()).Info("init redis")

	// New redis pool
	pool = &redis.Pool{
		MaxIdle:     config.MaxIdle,
		MaxActive:   config.MaxActive,
		Wait:        true,
		IdleTimeout: time.Duration(config.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
		},
	}
}
