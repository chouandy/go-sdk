package auditable

import (
	"database/sql/driver"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// Changes changes struct
type Changes struct {
	Was map[string]interface{} `json:"was"`
	Is  map[string]interface{} `json:"is"`
}

// Value for gorm query
func (c Changes) Value() (driver.Value, error) {
	return jsonex.Marshal(c)
}

// Scan for gorm query
func (c *Changes) Scan(value interface{}) error {
	switch value.(type) {
	case string:
		return jsonex.Unmarshal([]byte(value.(string)), c)
	case []byte:
		return jsonex.Unmarshal(value.([]byte), c)
	}

	return errors.New("unknown changes type")
}

// String changes to string
func (c *Changes) String() string {
	data, _ := jsonex.Marshal(c)
	return string(data)
}

// IsNothingChanged nothing changed
func (c *Changes) IsNothingChanged() bool {
	return len(c.Was) == 0 && len(c.Is) == 0
}

// NewChanges new changes
func NewChanges(db *gorm.DB, action uint32) (*Changes, error) {
	changes := &Changes{
		Was: make(map[string]interface{}),
		Is:  make(map[string]interface{}),
	}

	switch action {
	case ActionCreate:
		// Get auditable fields
		for _, f := range db.Statement.Schema.Fields {
			if IsAuditableField(f) {
				changes.Is[f.DBName] = f.ReflectValueOf(db.Statement.ReflectValue).Addr().Interface()
			}
		}
	case ActionUpdate:
		// Get original entity field
		originalEntityField := db.Statement.Schema.LookUpField(FieldNameOriginalEntity)
		if originalEntityField == nil {
			return nil, errors.New("original entity not found")
		}
		fmt.Println(fmt.Sprintf("originalEntityField %+v", originalEntityField))

		return nil, nil
		// New original entity scope
		// originalEntityScope := scope.NewDB().NewScope(
		// 	originalEntityField.Field.Elem().Interface(),
		// )
		// // Get auditable fields
		// for _, nf := range db.Statement.Schema.Fields {
		// 	if IsAuditableField(nf) {
		// 		// Get old field by new field name
		// 		if of, ok := originalEntityScope.FieldByName(nf.Name); ok {
		// 			// New old value and new value variables
		// 			var ov, nv interface{}
		// 			// Check filed is pointer or not
		// 			if nf.Field.Kind() == reflect.Ptr {
		// 				// of is blank ? nil : value
		// 				if of.IsBlank {
		// 					ov = nil
		// 				} else {
		// 					ov = of.Field.Elem().Interface()
		// 				}
		// 				// nf is blank ? nil : value
		// 				if nf.IsBlank {
		// 					nv = nil
		// 				} else {
		// 					nv = nf.Field.Elem().Interface()
		// 				}
		// 			} else {
		// 				ov = of.Field.Interface()
		// 				nv = nf.Field.Interface()

		// 			}
		// 			// Check ov and nv is the same or not
		// 			if ov != nv {
		// 				changes.Was[nf.DBName] = ov
		// 				changes.Is[nf.DBName] = nv
		// 			}
		// 		}
		// 	}
		// }
		// // Check is nothing changed or not
		// if changes.IsNothingChanged() {
		// 	return nil, errors.New("nothing changed")
		// }
	case ActionDelete:
		// Get auditable fields
		for _, f := range db.Statement.Schema.Fields {
			if IsAuditableField(f) {
				changes.Was[f.DBName] = f.ReflectValueOf(db.Statement.ReflectValue).Addr().Interface()
			}
		}
	}

	return changes, nil
}
