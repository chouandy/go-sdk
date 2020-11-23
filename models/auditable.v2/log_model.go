package auditable

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"gorm.io/gorm"
)

// LogModel auditable log model struct
type LogModel struct {
	ID          uint64    `json:"id"`
	TriggerID   uint64    `json:"-"`
	AuditableID uint64    `json:"-"`
	Action      uint32    `json:"action"`
	Changes     Changes   `json:"changes" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
}

// NewLog new log
func NewLog(db *gorm.DB, action uint32) (interface{}, error) {
	// Get log model type
	var modelType reflect.Type
	for _, f := range db.Statement.Schema.Fields {
		if IsAuditableLogsField(f) {
			modelType = db.Statement.Schema.ModelType
			break
		}
	}
	if modelType == nil {
		return nil, errors.New("auditable logs type not found")
	}

	// New changes
	changes, err := NewChanges(db, action)
	if err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("changes %+v", changes))

	// New log model instance
	tx := db.Session(&gorm.Session{NewDB: true})
	log := tx.Model(reflect.New(modelType).Interface())

	// Set columns
	if field := db.Statement.Schema.LookUpField(FieldNameTriggerID); field != nil {
		log.Set(FieldNameTriggerID, field.ReflectValueOf(db.Statement.ReflectValue).Addr().Interface())
	}

	primarykeyField := db.Statement.Schema.PrimaryFields[0]

	// fmt.Println(fmt.Sprintf("Schema %+v", db.Statement.Schema.LookUpField(db.Statement.Schema.PrimaryFields[0].ValueOf())))

	log.Statement.SetColumn(FieldNameAuditableID, primarykeyField.ReflectValueOf(db.Statement.ReflectValue).Addr().Interface())
	// newScope.SetColumn(FieldNameAuditableID, scope.PrimaryKeyValue())
	log.Statement.SetColumn(FieldNameAction, action)
	log.Statement.SetColumn(FieldNameChanges, changes)

	fmt.Println(fmt.Sprintf("log1 %+v", log))

	return log, nil
}
