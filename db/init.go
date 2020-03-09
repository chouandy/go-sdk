package db

import logex "github.com/chouandy/go-sdk/log"

// Init init
func Init() (err error) {
	// New config
	if err = NewConfig(); err != nil {
		return
	}

	// Print log
	logex.Log.WithFields(config.LogrusFields()).Info("init db")

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
