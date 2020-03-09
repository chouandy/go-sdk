package db

import (
	"database/sql"
)

// CreateDatabase create database
func CreateDatabase() error {
	// New db connection
	conn, err := sql.Open(config.GetDriver(), config.DataSourceWithoutDatabase())
	if err != nil {
		return err
	}
	// defer close db connection
	defer conn.Close()

	// Exec statement
	if _, err = conn.Exec(config.CreateDatabaseStatement()); err != nil {
		return err
	}

	return nil
}

// DropDatabase drop database
func DropDatabase() error {
	// New db connection
	conn, err := sql.Open(config.GetDriver(), config.DataSourceWithoutDatabase())
	if err != nil {
		return err
	}
	// defer close db connection
	defer conn.Close()

	// Exec statement
	if _, err = conn.Exec(config.DropDatabaseStatement()); err != nil {
		return err
	}

	return nil
}
