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
	assert.Nil(t, err)
	defer client.Close()
}
