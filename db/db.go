package db

import dgoex "github.com/chouandy/go-sdk/log"

var config *Config

// Init init
func Init() (err error) {
	// New config
	config, err = NewConfig()
	if err != nil {
		return
	}

	// Print log
	dgoex.TextLog().WithFields(config.LogrusFields()).Info("init db")

	// Init gorm db
	if err = InitGORMDB(); err != nil {
		return
	}

	// Init sqlx db
	if err = InitSQLXDB(); err != nil {
		return
	}

	// Load sql files
	if err = LoadSQLFiles(); err != nil {
		return
	}

	return
}
