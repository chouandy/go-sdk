package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"

	logex "github.com/chouandy/go-sdk/log"
)

var config *Config

var pool *redis.Pool

// Init init redis pool
func Init() (err error) {
	// New config
	config, err = NewConfig()
	if err != nil {
		return
	}

	// Print log
	logex.TextLog().WithFields(config.LogrusFields()).Info("init redis")

	// New redis pool
	pool = &redis.Pool{
		MaxIdle:     config.MaxIdle,
		MaxActive:   config.MaxActive,
		Wait:        true,
		IdleTimeout: time.Duration(config.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(config.URL)
		},
	}

	return nil
}
