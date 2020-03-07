package auditable

import (
	"database/sql/driver"
	"errors"
	"reflect"

	"github.com/jinzhu/gorm"
)

// Changes changes struct
type Changes struct {
	Was map[string]interface{} `json:"was"`
	Is  map[string]interface{} `json:"is"`
}

// Value for gorm query
func (c Changes) Value() (driver.Value, error) {
	data, err := jsonex.Marshal(c)
	return string(data), err
}

// Scan for gorm query
func (c *Changes) Scan(value interface{}) error {
	return jsonex.Unmarshal([]byte(value.(string)), c)
}

// String changes to string
func (c *Changes) String() string {
	data, _ := jsonex.Marshal(c)
	return string(data)
}

// IsNothingChanged nothing changed
func (c *Changes) IsNothingChanged() bool {
	return len(c.Was) == 0 || len(c.Is) == 0
}

// NewChanges new changes
func NewChanges(scope *gorm.Scope, action uint32) (*Changes, error) {
	changes := &Changes{
		Was: make(map[string]interface{}),
		Is:  make(map[string]interface{}),
	}

	switch action {
	case ActionCreate:
		// Get auditable fields
		for _, f := range scope.Fields() {
			if IsAuditableField(f) && !f.IsBlank {
				changes.Is[f.DBName] = f.Field.Interface()
			}
		}
	case ActionUpdate:
		// Get original entity field
		originalEntityField, ok := scope.FieldByName(FieldNameOriginalEntity)
		if !ok {
			return nil, errors.New("original entity not found")
		}
		// New original entity scope
		originalEntityScope := scope.NewDB().NewScope(
			originalEntityField.Field.Elem().Interface(),
		)
		// Get auditable fields
		for _, nf := range scope.Fields() {
			if IsAuditableField(nf) {
				// Get old field by new field name
				if of, ok := originalEntityScope.FieldByName(nf.Name); ok {
					// New old value and new value variables
					var ov, nv interface{}
					// Check filed is pointer or not
					if nf.Field.Kind() == reflect.Ptr {
						// of is blank ? nil : value
						if of.IsBlank {
							ov = nil
						} else {
							ov = of.Field.Elem().Interface()
						}
						// nf is blank ? nil : value
						if nf.IsBlank {
							nv = nil
						} else {
							nv = nf.Field.Elem().Interface()
						}
					} else {
						ov = of.Field.Interface()
						nv = nf.Field.Interface()
					}
					// Check ov and nv is the same or not
					if ov != nv {
						changes.Was[nf.DBName] = ov
						changes.Is[nf.DBName] = nv
					}
				}
			}
		}
		// Check is nothing changed or not
		if changes.IsNothingChanged() {
			return nil, errors.New("nothing changed")
		}
	case ActionDelete:
		// Get auditable fields
		for _, f := range scope.Fields() {
			if IsAuditableField(f) && !f.IsBlank {
				changes.Was[f.DBName] = f.Field.Interface()
			}
		}
	}

	return changes, nil
}
