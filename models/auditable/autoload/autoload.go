package autoload

import (
	"github.com/chouandy/go-sdk/models/auditable"
	"github.com/jinzhu/gorm"
)

func init() {
	gorm.DefaultCallback.Query().After("gorm:after_query").Register("auditable:query", auditable.QueryCallback)
	gorm.DefaultCallback.Create().After("gorm:after_create").Register("auditable:create", auditable.CreateCallback)
	gorm.DefaultCallback.Update().After("gorm:after_update").Register("auditable:update", auditable.UpdateCallback)
	gorm.DefaultCallback.Delete().After("gorm:after_delete").Register("auditable:delete", auditable.DeleteCallback)
}
