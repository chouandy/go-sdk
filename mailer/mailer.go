package mailer

import logex "github.com/chouandy/go-sdk/log"

var config *Config

// Init init
func Init() (err error) {
	// New config
	config, err = NewConfig()
	if err != nil {
		return
	}

	// Print log
	logex.TextLog().WithFields(config.LogrusFields()).Info("init mailer")

	// Load template files
	if err = LoadTemplateFiles(); err != nil {
		return
	}

	return nil
}

// URL url
func URL() string {
	return config.URL()
}
