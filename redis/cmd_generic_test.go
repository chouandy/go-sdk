package redis_test

import (
	. "github.com/chouandy/go-sdk/redis"

	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func TestDel(t *testing.T) {
	TestInit(t)

	client, err := NewClient()
	assert.Nil(t, err)
	defer client.Close()

	keys := []string{"a", "b", "c"}

	// Del before set
	err = client.Del(keys...)
	assert.Nil(t, err)

	// Set
	err = client.Set("a", "a", 0)
	assert.Nil(t, err)
	err = client.Set("b", "b", 0)
	assert.Nil(t, err)
	err = client.Set("c", "c", 0)
	assert.Nil(t, err)

	// Del
	err = client.Del(keys...)
	assert.Nil(t, err)

	// Check
	_, err = redis.String(client.Get("a"))
	assert.EqualError(t, err, "redigo: nil returned")
	_, err = redis.String(client.Get("b"))
	assert.EqualError(t, err, "redigo: nil returned")
	_, err = redis.String(client.Get("c"))
	assert.EqualError(t, err, "redigo: nil returned")
}

func TestScan(t *testing.T) {
	TestInit(t)

	client, err := NewClient()
	assert.Nil(t, err)
	defer client.Close()

	keys := []string{"1:a", "1:b", "1:c"}

	// Del before set
	err = client.Del(keys...)
	assert.Nil(t, err)

	// Set
	err = client.Set("1:a", "a", 0)
	assert.Nil(t, err)
	err = client.Set("1:b", "b", 0)
	assert.Nil(t, err)
	err = client.Set("1:c", "c", 0)
	assert.Nil(t, err)

	// Scan
	values, err := redis.Values(client.Scan(0, "1:*", 0))
	assert.Nil(t, err)

	// Get next cursor
	cursor, err := redis.Uint64(values[0], nil)
	assert.Nil(t, err)
	assert.Equal(t, uint64(0), cursor)

	// Get keys
	scannedKeys, err := redis.Strings(values[1], nil)
	assert.Nil(t, err)
	for _, key := range scannedKeys {
		assert.Contains(t, keys, key)
	}

	// Del
	err = client.Del(keys...)
	assert.Nil(t, err)

	// Check
	_, err = redis.String(client.Get("1:a"))
	assert.EqualError(t, err, "redigo: nil returned")
	_, err = redis.String(client.Get("1:b"))
	assert.EqualError(t, err, "redigo: nil returned")
	_, err = redis.String(client.Get("1:c"))
	assert.EqualError(t, err, "redigo: nil returned")
}
