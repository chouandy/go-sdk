package redis_test

import (
	. "github.com/chouandy/go-sdk/redis"
	"github.com/stretchr/testify/assert"

	"testing"
)

var testRedisKey = "redistest"

func TestInit(t *testing.T) {
	err := Init()
	assert.Nil(t, err)
}
