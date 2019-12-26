package redis_test

import (
	. "github.com/chouandy/go-sdk/redis"

	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

func TestHSetHGetHDel(t *testing.T) {
	TestInit(t)

	client, err := NewClient()
	assert.Nil(t, err)
	defer client.Close()

	values := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
		"d": "D",
		"e": "E",
	}

	for field, value := range values {
		err = client.HSet(testRedisKey, field, value)
		assert.Nil(t, err)
	}

	for field, value := range values {
		value2, err := redis.String(client.HGet(testRedisKey, field))
		assert.Nil(t, err)
		assert.Equal(t, value, value2)
	}

	client.HDel(testRedisKey, "a")
	_, err = redis.String(client.HGet(testRedisKey, "a"))
	assert.EqualError(t, err, "redigo: nil returned")

	client.HDel(testRedisKey, "b")
	_, err = redis.String(client.HGet(testRedisKey, "b"))
	assert.EqualError(t, err, "redigo: nil returned")

	err = client.Del(testRedisKey)
	assert.Nil(t, err)
}
