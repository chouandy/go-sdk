package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// Official Document: https://redis.io/commands#string

// Set https://redis.io/commands/set
func (c *Client) Set(key string, value interface{}, expiration time.Duration) (err error) {
	args := redis.Args{key, value}
	if expiration > 0 {
		if usePrecise(expiration) {
			args = args.Add("PX").Add(formatMs(expiration))
		} else {
			args = args.Add("EX").Add(formatSec(expiration))
		}
	}

	_, err = c.Do("SET", args...)
	return
}

// Get https://redis.io/commands/get
func (c *Client) Get(key string) (interface{}, error) {
	return c.Do("GET", key)
}

// SetWithRemainingTTL set with remaining ttl
func (c *Client) SetWithRemainingTTL(key string, value interface{}) (err error) {
	// New script
	script := redis.NewScript(0,
		"local ttl = redis.call('ttl', ARGV[1]) if ttl > 0 then return redis.call('SETEX', ARGV[1], ttl, ARGV[2]) end",
	)

	// Execute script
	_, err = script.Do(c.Conn, key, value)
	return
}
