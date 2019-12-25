package redis_test

import (
	. "github.com/chouandy/go-sdk/redis"

	"testing"

	configex "github.com/chouandy/go-sdk/config"
)

var testRedisKey = "redistest"

type testConfig struct {
	Redis Config
}

func TestInit(t *testing.T) {
	// Load config
	configex.SetConfigDir("../test/config")
	config := new(testConfig)
	configex.Load(config, false)

	// Init redis
	Init(config.Redis)
}
