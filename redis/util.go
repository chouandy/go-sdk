package redis

import (
	"time"

	"github.com/chouandy/go-sdk/log"
	"github.com/gomodule/redigo/redis"
)

func usePrecise(dur time.Duration) bool {
	return dur < time.Second || dur%time.Second != 0
}

func formatMs(dur time.Duration) int64 {
	if dur > 0 && dur < time.Millisecond {
		log.Log.Errorf(
			"specified duration is %s, but minimal supported value is %s",
			dur, time.Millisecond,
		)
	}
	return int64(dur / time.Millisecond)
}

func formatSec(dur time.Duration) int64 {
	if dur > 0 && dur < time.Second {
		log.Log.Errorf(
			"specified duration is %s, but minimal supported value is %s",
			dur, time.Second,
		)
	}
	return int64(dur / time.Second)
}

// IsErrNil is err nil
func IsErrNil(err error) bool {
	return err == redis.ErrNil
}
