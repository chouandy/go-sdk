package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// Official Document: https://redis.io/commands#generic

// Del https://redis.io/commands/del
func (c *Client) Del(keys ...string) (err error) {
	args := redis.Args{}.AddFlat(keys)
	_, err = c.Do("DEL", args...)
	return
}

// Scan https://redis.io/commands/scan
func (c *Client) Scan(cursor uint64, match string, count int64) (interface{}, error) {
	args := redis.Args{cursor}

	if len(match) > 0 {
		args = args.Add("MATCH").Add(match)
	}

	if count > 0 {
		args = args.Add("COUNT").Add(count)
	}

	return c.Do("SCAN", args...)
}

// Expire https://redis.io/commands/expire
func (c *Client) Expire(key string, expiration time.Duration) (err error) {
	args := redis.Args{key, formatSec(expiration)}
	_, err = c.Do("EXPIRE", args...)
	return
}

// ExpireAt https://redis.io/commands/expireat
func (c *Client) ExpireAt(key string, tm time.Time) (err error) {
	args := redis.Args{key, tm.Unix()}
	_, err = c.Do("EXPIREAT", args...)
	return
}
