package redis_test

import (
	. "github.com/chouandy/go-sdk/redis"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	_, err := NewClient()
	assert.EqualError(t, err, "Init redis pool first")

	TestInit(t)

	client, err := NewClient()
	defer client.Close()
	assert.Nil(t, err)
}
