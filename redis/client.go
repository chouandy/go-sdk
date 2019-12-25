package redis

import (
	"errors"

	"github.com/gomodule/redigo/redis"
)

// Client client struct
type Client struct {
	redis.Conn
}

// NewClient new client
func NewClient() (*Client, error) {
	// Check pool
	if pool == nil {
		return nil, errors.New("Init redis pool first")
	}

	// Get connection
	conn := pool.Get()

	return &Client{Conn: conn}, nil
}
