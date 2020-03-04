package autoload

import dbex "github.com/chouandy/go-sdk/db"

func init() {
	if err := dbex.Init(); err != nil {
		panic(err)
	}
}
