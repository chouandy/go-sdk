package db_test

import (
	. "github.com/chouandy/go-sdk/db"

	"fmt"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestInitV2(t *testing.T) {
	// Set test cases
	testCases := []struct {
		databaseURL string
	}{
		{
			databaseURL: os.Getenv("MYSQL_DATABASE_URL"),
		},
		{
			databaseURL: os.Getenv("POSTGRES_DATABASE_URL"),
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			os.Setenv("DATABASE_URL", testCase.databaseURL)

			var err error
			err = NewConfig()
			assert.Nil(t, err)
			err = CreateDatabase()
			assert.Nil(t, err)

			// Init db
			err = InitV2()
			assert.Nil(t, err)

			db1 := GORMV2()
			assert.IsType(t, &gorm.DB{}, db1)
			sqlDB, err := GORMV2().DB()
			assert.Nil(t, err)
			err = sqlDB.Close()
			assert.Nil(t, err)

			db2 := SQLX()
			assert.IsType(t, &sqlx.DB{}, db2)
			err = SQLX().Close()
			assert.Nil(t, err)

			err = DropDatabase()
			assert.Nil(t, err)
		})
	}

	os.Setenv("DATABASE_URL", "")
}
