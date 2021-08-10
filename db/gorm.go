package db

import (
	"sync/atomic"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

func NewGORMDialector() (dialector gorm.Dialector, err error) {
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

// NewGORMDB new gorm db
func NewGORMDB() (*gorm.DB, error) {
	dialector, err := NewGORMDialector()
	if err != nil {
		return nil, err
	}

	// New db
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

// GORM get gorm db
func GORM() *gorm.DB {
	// Get db
	v := gormDB.Load()

	// Convert db type
	db := v.(*gorm.DB)

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
		if newDB, err := NewGORMDB(); err == nil {
			// Close old db
			if sqlDB != nil {
				sqlDB.Close()
			}
			// Set new db
			gormDB.Store(newDB)

			return newDB
		}
	}

	return db
}
