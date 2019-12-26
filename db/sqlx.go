package db

import (
	"sync/atomic"
	"time"

	"github.com/jmoiron/sqlx"
)

var sqlxDB atomic.Value

// InitSQLXDB init sqlx db
func InitSQLXDB() error {
	// New sqlx db
	db, err := NewSQLXDB()
	if err != nil {
		return err
	}

	// Set sqlx db
	sqlxDB.Store(db)

	return nil
}

// NewSQLXDB new sqlx db
func NewSQLXDB() (*sqlx.DB, error) {
	// New db
	db, err := sqlx.Connect(config.Driver, config.DataSource())
	if err != nil {
		return nil, err
	}

	// Setup db
	db.DB.SetMaxOpenConns(config.MaxOpenConns)
	db.DB.SetMaxIdleConns(config.MaxIdleConns)
	db.DB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

// SQLX get sqlx db
func SQLX() *sqlx.DB {
	// Get db
	v := sqlxDB.Load()

	// Convert db type
	db := v.(*sqlx.DB)

	// Check db
	if err := db.DB.Ping(); err != nil {
		// New db
		if newDB, err := NewSQLXDB(); err == nil {
			// Close old db
			db.Close()
			// Set new db
			sqlxDB.Store(newDB)

			return newDB
		}
	}

	return db
}
