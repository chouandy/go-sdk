package redis_test

import (
	. "github.com/chouandy/go-sdk/redis"

	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func TestSetGet(t *testing.T) {
	TestInit(t)

	client, err := NewClient()
	defer client.Close()
	assert.Nil(t, err)

	// Without expiration
	// - Set
	err = client.Set("a", "a", 0)
	assert.Nil(t, err)
	// - Get
	a, err := redis.String(client.Get("a"))
	assert.Nil(t, err)
	assert.Equal(t, "a", a)
	// - Del
	err = client.Del("a")
	assert.Nil(t, err)
	// - Get after Del
	a, err = redis.String(client.Get("a"))
	assert.EqualError(t, err, "redigo: nil returned")

	// With millisecond expiration
	// - Set
	err = client.Set("b", "b", 1234*time.Millisecond)
	assert.Nil(t, err)
	// - Get
	b, err := redis.String(client.Get("b"))
	assert.Nil(t, err)
	assert.Equal(t, "b", b)
	// - Get expired
	time.Sleep(1500 * time.Millisecond)
	b, err = redis.String(client.Get("b"))
	assert.EqualError(t, err, "redigo: nil returned")

	// With second expiration
	// - Set
	err = client.Set("c", "c", 2*time.Second)
	assert.Nil(t, err)
	// - Get
	c, err := redis.String(client.Get("c"))
	assert.Nil(t, err)
	assert.Equal(t, "c", c)
	// - Get expired
	time.Sleep(2100 * time.Millisecond)
	c, err = redis.String(client.Get("c"))
	assert.EqualError(t, err, "redigo: nil returned")
}
