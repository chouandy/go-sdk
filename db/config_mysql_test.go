package db_test

import (
	. "github.com/chouandy/go-sdk/db"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMySQLConfigFromDatabaseURL(t *testing.T) {
	_, err := NewMySQLConfigFromDatabaseURL()
	assert.Equal(t, "database can't be blank", err.Error())

	os.Setenv("DATABASE_URL", "mysql://user:pass@tcp(host:3306)/db?charset=utf8&collation=utf8_general_ci")
	dialect, err := NewMySQLConfigFromDatabaseURL()
	assert.Nil(t, err)
	assert.Equal(t, "mysql", dialect.GetDriver())
	assert.Equal(t, "mysql://user:pass@tcp(host:3306)/db?charset=utf8&parseTime=true", dialect.DatabaseURL())
	assert.Equal(t, "user:pass@tcp(host:3306)/db?charset=utf8&parseTime=true", dialect.DataSource())
	assert.Equal(t, "user:pass@tcp(host:3306)/?charset=utf8", dialect.DataSourceWithoutDatabase())
	assert.Equal(t,
		"CREATE DATABASE `db` DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';",
		dialect.CreateDatabaseStatement(),
	)
	assert.Equal(t, "DROP DATABASE `db`", dialect.DropDatabaseStatement())

	os.Setenv("DATABASE_URL", "")
}
