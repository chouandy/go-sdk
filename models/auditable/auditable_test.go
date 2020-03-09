package auditable_test

import (
	. "github.com/chouandy/go-sdk/models/auditable"

	_ "github.com/chouandy/go-sdk/models/auditable/autoload"
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

	Logs []UserLog `auditable:"logs"`
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

			err = dbex.InitGORMDB()
			assert.Nil(t, err)

			err = dbex.GORM().AutoMigrate(&User{}, &UserLog{}).Error
			assert.Nil(t, err)

			var admin User
			var user User
			var userLogs []UserLog

			// -----------------------------------
			// -          Create Action          -
			// -----------------------------------

			admin = User{Email: "admin@example.com", Role: 1}
			err = dbex.GORM().Create(&admin).Error
			assert.Nil(t, err)

			user = User{Email: "user@example.com", Role: 0}
			user.SetTriggerID(admin.ID)
			err = dbex.GORM().Create(&user).Error
			assert.Nil(t, err)

			err = dbex.GORM().Preload("Trigger").Preload("Auditable").Find(&userLogs).Error
			assert.Nil(t, err)

			assert.Equal(t, uint64(1), userLogs[0].Trigger.ID)
			assert.Equal(t, uint64(2), userLogs[0].Auditable.ID)
			assert.Equal(t, uint32(0), userLogs[0].Auditable.Role)
			assert.Equal(t, `{"was":{},"is":{"role":0}}`, userLogs[0].Changes.String())

			// -----------------------------------
			// -          Update Action          -
			// -----------------------------------

			err = dbex.GORM().First(&user, user.ID).Error
			assert.Nil(t, err)
			user.SetTriggerID(admin.ID)
			user.SetOriginalEntity(user)

			updates := map[string]interface{}{
				"role": 1,
			}
			err = dbex.GORM().Model(&user).Updates(updates).Error
			assert.Nil(t, err)
			err = dbex.GORM().Preload("Trigger").Preload("Auditable").Order("id DESC").Find(&userLogs).Error
			assert.Nil(t, err)

			assert.Equal(t, uint64(1), userLogs[0].Trigger.ID)
			assert.Equal(t, uint64(2), userLogs[0].Auditable.ID)
			assert.Equal(t, uint32(1), userLogs[0].Auditable.Role)
			assert.Equal(t, `{"was":{"role":0},"is":{"role":1}}`, userLogs[0].Changes.String())

			err = dbex.GORM().Close()
			assert.Nil(t, err)

			err = dbex.DropDatabase()
			assert.Nil(t, err)
		})
	}

	os.Setenv("DATABASE_URL", "")
}
