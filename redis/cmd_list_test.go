package redis_test

import (
	. "github.com/chouandy/go-sdk/redis"

	"encoding/json"
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestLPushLPop(t *testing.T) {
	TestInit(t)

	client, err := NewClient()
	assert.Nil(t, err)
	defer client.Close()

	values := []string{"a", "b", "c", "d", "e"}

	for _, value := range values {
		err = client.LPush(testRedisKey, value)
		assert.Nil(t, err)
	}

	values2, err := redis.Strings(client.LRange(testRedisKey, 0, -1))
	assert.Nil(t, err)
	assert.Equal(t, []string{"e", "d", "c", "b", "a"}, values2)

	value, err := redis.String(client.LPop(testRedisKey))
	assert.Nil(t, err)
	assert.Equal(t, "e", value)

	value, err = redis.String(client.LPop(testRedisKey))
	assert.Nil(t, err)
	assert.Equal(t, "d", value)

	value, err = redis.String(client.LPop(testRedisKey))
	assert.Nil(t, err)
	assert.Equal(t, "c", value)

	values3, err := redis.Strings(client.LRange(testRedisKey, 0, -1))
	assert.Nil(t, err)
	assert.Equal(t, []string{"b", "a"}, values3)

	err = client.Del(testRedisKey)
	assert.Nil(t, err)
}

func TestRPushRPop(t *testing.T) {
	TestInit(t)

	client, err := NewClient()
	assert.Nil(t, err)
	defer client.Close()

	values := []string{"a", "b", "c", "d", "e"}

	for _, value := range values {
		err = client.RPush(testRedisKey, value)
		assert.Nil(t, err)
	}

	values2, err := redis.Strings(client.LRange(testRedisKey, 0, -1))
	assert.Nil(t, err)
	assert.Equal(t, []string{"a", "b", "c", "d", "e"}, values2)

	value, err := redis.String(client.RPop(testRedisKey))
	assert.Nil(t, err)
	assert.Equal(t, "e", value)

	value, err = redis.String(client.RPop(testRedisKey))
	assert.Nil(t, err)
	assert.Equal(t, "d", value)

	value, err = redis.String(client.RPop(testRedisKey))
	assert.Nil(t, err)
	assert.Equal(t, "c", value)

	values3, err := redis.Strings(client.LRange(testRedisKey, 0, -1))
	assert.Nil(t, err)
	assert.Equal(t, []string{"a", "b"}, values3)

	err = client.Del(testRedisKey)
	assert.Nil(t, err)
}

func TestLPushLPopStruct(t *testing.T) {
	TestInit(t)

	client, err := NewClient()
	assert.Nil(t, err)
	defer client.Close()

	type Order struct {
		ID     string          `json:"order_id"`
		Amount decimal.Decimal `json:"amount"`
		IsBuy  bool            `json:"is_buy"`
		Status int32           `json:"status"`
	}

	orders := []Order{
		{
			ID:     "122162",
			Amount: decimal.NewFromFloat(1.2345),
			IsBuy:  false,
			Status: 0,
		},
		{
			ID:     "122163",
			Amount: decimal.NewFromFloat(2.3456),
			IsBuy:  true,
			Status: 0,
		},
		{
			ID:     "122164",
			Amount: decimal.NewFromFloat(3.4567),
			IsBuy:  true,
			Status: 0,
		},
	}

	// [2, 1, 0]
	for _, order := range orders {
		data, err := json.Marshal(order)
		assert.Nil(t, err)

		err = client.LPush(testRedisKey, data)
		assert.Nil(t, err)
	}

	var order Order

	value, err := redis.Bytes(client.LPop(testRedisKey))
	assert.Nil(t, err)
	err = json.Unmarshal(value, &order)
	assert.Nil(t, err)
	assert.Equal(t, orders[2], order)

	value, err = redis.Bytes(client.RPop(testRedisKey))
	assert.Nil(t, err)
	err = json.Unmarshal(value, &order)
	assert.Nil(t, err)
	assert.Equal(t, orders[0], order)

	err = client.Del(testRedisKey)
	assert.Nil(t, err)
}
