package auditable

import (
	"gorm.io/gorm"
)

// Init init
func Init(db *gorm.DB) {
	db.Callback().Create().After("gorm:create").Register("auditable:create", CreateCallback)
	db.Callback().Update().After("gorm:update").Register("auditable:update", UpdateCallback)
	db.Callback().Delete().After("gorm:delete").Register("auditable:delete", DeleteCallback)
}
