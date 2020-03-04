package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMySQLConfigFromDatabaseURL(t *testing.T) {
	config, err := NewMySQLConfigFromDatabaseURL()
	assert.Equal(t, "database can't be blank", err.Error())

	os.Setenv("DATABASE_URL", "mysql://user:pass@tcp(host:3306)/db?collation=utf8mb4_general_ci")
	config, err = NewMySQLConfigFromDatabaseURL()
	assert.Nil(t, err)
	assert.Equal(t, "mysql", config.GetDriver())
	assert.Equal(t, "mysql://user:pass@tcp(host:3306)/db?charset=utf8&parseTime=true", config.DatabaseURL())
	assert.Equal(t, "user:pass@tcp(host:3306)/db?charset=utf8&parseTime=true", config.DataSource())
	assert.Equal(t, "user:pass@tcp(host:3306)/?charset=utf8", config.DataSourceWithoutDatabase())
	assert.Equal(t,
		"CREATE DATABASE `db` DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8mb4_general_ci';",
		config.CreateDatabaseStatement(),
	)
	assert.Equal(t, "DROP DATABASE `db`", config.DropDatabaseStatement())

	os.Setenv("DATABASE_URL", "")
}
