package db

import (
	"sync/atomic"
	"time"

	"github.com/jinzhu/gorm"
)

var gormDB atomic.Value

// InitGORMDB init gorm db
func InitGORMDB() error {
	// New gorm db
	db, err := NewGORMDB()
	if err != nil {
		return err
	}

	// Set gorm db
	gormDB.Store(db)

	return nil
}

// NewGORMDB new gorm db
func NewGORMDB() (*gorm.DB, error) {
	// New db
	db, err := gorm.Open(config.Driver, config.DataSource())
	if err != nil {
		return nil, err
	}

	// Setup db
	db.LogMode(config.LogMode)
	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Hour)

	return db, nil
}

// GORM get gorm db
func GORM() *gorm.DB {
	// Get db
	v := gormDB.Load()

	// Convert db type
	db := v.(*gorm.DB)

	// Check db
	if err := db.DB().Ping(); err != nil {
		// New db
		if newDB, err := NewGORMDB(); err == nil {
			// Close old db
			db.Close()
			// Set new db
			gormDB.Store(newDB)

			return newDB
		}
	}

	return db
}
