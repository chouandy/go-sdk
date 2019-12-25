package redis

import "github.com/gomodule/redigo/redis"

// Official Document: https://redis.io/commands#list

// LPush https://redis.io/commands/lpush
func (c *Client) LPush(key string, values ...interface{}) (err error) {
	args := redis.Args{key}.Add(values...)
	_, err = c.Do("LPUSH", args...)
	return
}

// LPop https://redis.io/commands/lpop
func (c *Client) LPop(key string) (interface{}, error) {
	return c.Do("LPOP", key)
}

// LRange https://redis.io/commands/lrange
func (c *Client) LRange(key string, start, stop int64) (interface{}, error) {
	return c.Do("LRANGE", key, start, stop)
}

// RPush https://redis.io/commands/rpush
func (c *Client) RPush(key string, values ...interface{}) (err error) {
	args := redis.Args{key}.Add(values...)
	_, err = c.Do("RPUSH", args...)
	return
}

// RPop https://redis.io/commands/rpop
func (c *Client) RPop(key string) (interface{}, error) {
	return c.Do("RPOP", key)
}
