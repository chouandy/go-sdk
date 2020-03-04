package autoload

import mailerex "github.com/chouandy/go-sdk/mailer"

func init() {
	if err := mailerex.Init(); err != nil {
		panic(err)
	}
}
