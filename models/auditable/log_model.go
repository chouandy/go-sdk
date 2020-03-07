package auditable

import (
	"errors"
	"reflect"
	"time"

	"github.com/jinzhu/gorm"
)

// LogModel auditable log model struct
type LogModel struct {
	ID          uint64    `json:"id"`
	TriggerID   uint64    `json:"-"`
	AuditableID uint64    `json:"-"`
	Action      uint32    `json:"action"`
	Changes     Changes   `json:"changes"`
	CreatedAt   time.Time `json:"created_at"`
}

// NewLog new log
func NewLog(scope *gorm.Scope, action uint32) (interface{}, error) {
	// Get log model type
	var logModelType reflect.Type
	for _, f := range scope.Fields() {
		if IsAuditableLogsField(f) {
			logModelType = f.Struct.Type.Elem()
			break
		}
	}
	if logModelType == nil {
		return nil, errors.New("auditable logs type not found")
	}

	// New changes
	changes, err := NewChanges(scope, action)
	if err != nil {
		return nil, err
	}

	// New log model instance
	newDB := scope.NewDB()
	log := reflect.New(logModelType).Interface()
	newScope := newDB.NewScope(log)

	// Set columns
	if f, ok := scope.FieldByName(FieldNameTriggerID); ok {
		newScope.SetColumn(FieldNameTriggerID, f.Field.Elem().Interface())
	}
	newScope.SetColumn(FieldNameAuditableID, scope.PrimaryKeyValue())
	newScope.SetColumn(FieldNameAction, action)
	newScope.SetColumn(FieldNameChanges, changes)

	return log, nil
}
