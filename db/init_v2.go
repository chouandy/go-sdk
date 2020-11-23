package db

import logex "github.com/chouandy/go-sdk/log"

// InitV2 init v2
func InitV2() (err error) {
	// New config
	if err = NewConfig(); err != nil {
		return
	}

	// Print log
	logex.Log.WithFields(config.LogrusFields()).Info("init db")

	// Init gorm v2 db
	if err = InitGORMV2DB(); err != nil {
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
