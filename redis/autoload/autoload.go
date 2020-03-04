package autoload

import redisex "github.com/chouandy/go-sdk/redis"

func init() {
	if err := redisex.Init(); err != nil {
		panic(err)
	}
}
