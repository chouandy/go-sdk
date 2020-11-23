package autoload

import dbex "github.com/chouandy/go-sdk/db"

func init() {
	if err := dbex.InitV2(); err != nil {
		panic(err)
	}
}
