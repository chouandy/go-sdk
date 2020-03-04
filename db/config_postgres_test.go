package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPostgresConfigFromDatabaseURL(t *testing.T) {
	config, err := NewPostgresConfigFromDatabaseURL()
	assert.Equal(t, "invalid connection protocol: ", err.Error())

	os.Setenv("DATABASE_URL", "postgres://user:pass@host:5432/db?sslmode=disable")
	config, err = NewPostgresConfigFromDatabaseURL()
	assert.Nil(t, err)
	assert.Equal(t, "postgres", config.GetDriver())
	assert.Equal(t, "postgres://user:pass@host:5432/db?sslmode=disable", config.DatabaseURL())
	assert.Equal(t, "postgres://user:pass@host:5432/db?sslmode=disable", config.DataSource())
	assert.Equal(t, "postgres://user:pass@host:5432/?sslmode=disable", config.DataSourceWithoutDatabase())
	assert.Equal(t,
		`CREATE DATABASE "db" ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8';`,
		config.CreateDatabaseStatement(),
	)
	assert.Equal(t, `DROP DATABASE "db"`, config.DropDatabaseStatement())

	os.Setenv("DATABASE_URL", "")
}
