package db_test

import (
	. "github.com/chouandy/go-sdk/db"
	_ "github.com/go-sql-driver/mysql"

	"testing"

	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	// Init db
	err := Init()
	assert.Nil(t, err)

	db1 := GORM()
	assert.IsType(t, &gorm.DB{}, db1)

	db2 := SQLX()
	assert.IsType(t, &sqlx.DB{}, db2)
}
