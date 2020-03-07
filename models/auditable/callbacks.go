package auditable

import (
	"github.com/jinzhu/gorm"
)

// QueryCallback query callback
func QueryCallback(scope *gorm.Scope) {
	// Check is auditable or not
	if !IsAuditableModel(scope) {
		return
	}

	// Set original entity
	if f, ok := scope.FieldByName(FieldNameOriginalEntity); ok {
		f.Set(scope.IndirectValue())
	}
}

// CreateCallback create callback
func CreateCallback(scope *gorm.Scope) {
	// Check is auditable and audit or not
	if !IsAuditableModel(scope) || !IsAudit(scope) {
		return
	}

	// New log
	log, err := NewLog(scope, ActionCreate)
	if err != nil {
		return
	}

	// Save log
	scope.Err(scope.NewDB().Save(log).Error)
}

// UpdateCallback update callback
func UpdateCallback(scope *gorm.Scope) {
	// Check is auditable and audit or not
	if !IsAuditableModel(scope) || !IsAudit(scope) {
		return
	}

	// New log
	log, err := NewLog(scope, ActionUpdate)
	if err != nil {
		return
	}

	// Save log
	scope.Err(scope.NewDB().Save(log).Error)
}

// DeleteCallback delete callback
func DeleteCallback(scope *gorm.Scope) {
	// Check is auditable and audit or not
	if !IsAuditableModel(scope) || !IsAudit(scope) {
		return
	}

	// New log
	log, err := NewLog(scope, ActionDelete)
	if err != nil {
		return
	}

	// Save log
	scope.Err(scope.NewDB().Save(log).Error)
}
