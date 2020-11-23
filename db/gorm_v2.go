package db

import (
	"sync/atomic"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormV2DB atomic.Value

// InitGORMV2DB init gorm v2 db
func InitGORMV2DB() error {
	// New gorm db
	db, err := NewGORMV2DB()
	if err != nil {
		return err
	}

	// Set gorm db
	gormV2DB.Store(db)

	return nil
}

// NewGORMV2DB new gorm db
func NewGORMV2DB() (*gorm.DB, error) {
	// new dialector
	dialector, err := NewGORMV2Dialector()
	if err != nil {
		return nil, err
	}

	// new db
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// debug mode
	if GetLogModeFromEnv() {
		db.Logger.LogMode(logger.Info)
	}

	// setup db
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(GetMaxOpenConnsFromEnv())
	sqlDB.SetMaxIdleConns(GetMaxIdleConnsFromEnv())
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

// GORMV2 get gorm v2 db
func GORMV2() *gorm.DB {
	// Get db
	v := gormV2DB.Load()

	// Convert db type
	db := v.(*gorm.DB)

	// Check db
	var shouldNewDB bool
	sqlDB, err := db.DB()
	if err != nil {
		shouldNewDB = true
	} else {
		if err := sqlDB.Ping(); err != nil {
			shouldNewDB = true
		}
	}

	if shouldNewDB {
		// New db
		if newDB, err := NewGORMV2DB(); err == nil {
			// Close old db
			sqlDB.Close()
			// Set new db
			gormDB.Store(newDB)

			return newDB
		}
	}

	return db
}

// NewGORMV2Dialector new gorm v2 dialector
func NewGORMV2Dialector() (dialector gorm.Dialector, err error) {
	driver := config.GetDriver()
	if driver == "mysql" {
		dialector = mysql.Open(config.DataSource())
	} else if driver == "postgres" {
		dialector = postgres.Open(config.DataSource())
	} else {
		err = errUnknownDBDriver
	}

	return
}
