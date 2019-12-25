package redis

import "github.com/gomodule/redigo/redis"

// Official Document: https://redis.io/commands#hash

// HSet https://redis.io/commands/hset
func (c *Client) HSet(key, field string, value interface{}) (err error) {
	_, err = c.Do("HSET", key, field, value)
	return
}

// HGet https://redis.io/commands/hget
func (c *Client) HGet(key, field string) (interface{}, error) {
	return c.Do("HGET", key, field)
}

// HDel https://redis.io/commands/hdel
func (c *Client) HDel(key string, fields ...string) (err error) {
	args := redis.Args{key}.AddFlat(fields)
	_, err = c.Do("HDEL", args...)
	return
}

// HKeys https://redis.io/commands/hkeys
func (c *Client) HKeys(key string) (interface{}, error) {
	return c.Do("HKEYS", key)
}

// HExists https://redis.io/commands/hexists
func (c *Client) HExists(key, field string) (bool, error) {
	return redis.Bool(c.Do("HEXISTS", key, field))
}
