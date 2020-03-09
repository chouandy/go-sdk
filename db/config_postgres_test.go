package db_test

import (
	. "github.com/chouandy/go-sdk/db"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPostgresConfigFromDatabaseURL(t *testing.T) {
	dialect, err := NewPostgresConfigFromDatabaseURL()
	assert.Equal(t, "invalid connection protocol: ", err.Error())

	os.Setenv("DATABASE_URL", "postgres://user:pass@host:5432/db?sslmode=disable")
	dialect, err = NewPostgresConfigFromDatabaseURL()
	assert.Nil(t, err)
	assert.Equal(t, "postgres", dialect.GetDriver())
	assert.Equal(t, "postgres://user:pass@host:5432/db?sslmode=disable", dialect.DatabaseURL())
	assert.Equal(t, "postgres://user:pass@host:5432/db?sslmode=disable", dialect.DataSource())
	assert.Equal(t, "postgres://user:pass@host:5432/?sslmode=disable", dialect.DataSourceWithoutDatabase())
	assert.Equal(t,
		`CREATE DATABASE "db" ENCODING = 'UTF8';`,
		dialect.CreateDatabaseStatement(),
	)
	assert.Equal(t, `DROP DATABASE "db"`, dialect.DropDatabaseStatement())

	os.Setenv("DATABASE_URL", "")
}
