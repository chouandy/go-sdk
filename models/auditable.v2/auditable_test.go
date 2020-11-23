// +build integration

package auditable_test

import (
	. "github.com/chouandy/go-sdk/models/auditable.v2"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	dbex "github.com/chouandy/go-sdk/db"
)

type User struct {
	Model

	ID    uint64
	Email string
	Role  uint32 `auditable:"field"`

	Logs []UserLog `auditable:"logs" gorm:"foreignKey:ID"`
}

type UserLog struct {
	LogModel

	Trigger   *User
	Auditable *User
}

func TestAuditable(t *testing.T) {
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

			// --------------------------
			// -          Init          -
			// --------------------------

			var err error
			err = dbex.NewConfig()
			assert.Nil(t, err)
			err = dbex.CreateDatabase()
			assert.Nil(t, err)

			err = dbex.InitGORMV2DB()
			assert.Nil(t, err)

			Init(dbex.GORMV2())

			err = dbex.GORMV2().AutoMigrate(&User{}, &UserLog{})
			assert.Nil(t, err)

			var admin User
			var user User
			var userLogs []UserLog

			// -----------------------------------
			// -          Create Action          -
			// -----------------------------------

			admin = User{Email: "admin@example.com", Role: 1}
			err = dbex.GORMV2().Create(&admin).Error
			assert.Nil(t, err)

			user = User{Email: "user@example.com", Role: 0}
			user.SetTriggerID(admin.ID)
			err = dbex.GORMV2().Create(&user).Error
			assert.Nil(t, err)

			err = dbex.GORMV2().Preload("Trigger").Preload("Auditable").Find(&userLogs).Error
			assert.Nil(t, err)

			assert.Equal(t, uint64(1), userLogs[0].Trigger.ID)
			assert.Equal(t, uint64(2), userLogs[0].Auditable.ID)
			assert.Equal(t, uint32(0), userLogs[0].Auditable.Role)
			assert.Equal(t, `{"was":{},"is":{"role":0}}`, userLogs[0].Changes.String())

			// 	// -----------------------------------
			// 	// -          Update Action          -
			// 	// -----------------------------------

			// 	err = dbex.GORMV2().First(&user, user.ID).Error
			// 	assert.Nil(t, err)
			// 	user.SetTriggerID(admin.ID)
			// 	user.SetOriginalEntity(user)

			// 	updates := map[string]interface{}{
			// 		"role": 1,
			// 	}
			// 	err = dbex.GORMV2().Model(&user).Updates(updates).Error
			// 	assert.Nil(t, err)
			// 	err = dbex.GORMV2().Preload("Trigger").Preload("Auditable").Order("id DESC").Find(&userLogs).Error
			// 	assert.Nil(t, err)

			// 	assert.Equal(t, uint64(1), userLogs[0].Trigger.ID)
			// 	assert.Equal(t, uint64(2), userLogs[0].Auditable.ID)
			// 	assert.Equal(t, uint32(1), userLogs[0].Auditable.Role)
			// 	assert.Equal(t, `{"was":{"role":0},"is":{"role":1}}`, userLogs[0].Changes.String())

			sqlDB, err := dbex.GORMV2().DB()
			assert.Nil(t, err)
			err = sqlDB.Close()
			assert.Nil(t, err)

			err = dbex.DropDatabase()
			assert.Nil(t, err)
		})
	}

	os.Setenv("DATABASE_URL", "")
}
