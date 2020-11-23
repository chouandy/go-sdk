package auditable

import (
	"fmt"

	"gorm.io/gorm"
)

// CreateCallback create callback
func CreateCallback(db *gorm.DB) {
	// Check is auditable and audit or not
	if !IsAuditableModel(db) || !IsAudit(db) {
		return
	}

	// New log
	log, err := NewLog(db, ActionCreate)
	if err != nil {
		return
	}

	fmt.Println(fmt.Sprintf("log %+v", log))

	// Save log
	db.Statement.AddError(db.Session(&gorm.Session{NewDB: true}).Save(log).Error)
}

// UpdateCallback update callback
func UpdateCallback(db *gorm.DB) {
	// Check is auditable and audit or not
	if !IsAuditableModel(db) || !IsAudit(db) {
		return
	}

	// New log
	log, err := NewLog(db, ActionUpdate)
	if err != nil {
		return
	}

	// Save log
	db.Statement.AddError(db.Session(&gorm.Session{NewDB: true}).Save(log).Error)
}

// DeleteCallback delete callback
func DeleteCallback(db *gorm.DB) {
	// Check is auditable and audit or not
	if !IsAuditableModel(db) || !IsAudit(db) {
		return
	}

	// New log
	log, err := NewLog(db, ActionDelete)
	if err != nil {
		return
	}

	// Save log
	db.Statement.AddError(db.Session(&gorm.Session{NewDB: true}).Save(log).Error)
}
