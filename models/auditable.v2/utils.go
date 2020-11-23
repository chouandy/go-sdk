package auditable

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// IsAuditableModel is auditable model
func IsAuditableModel(db *gorm.DB) bool {
	_, ok := db.Statement.Model.(Interface)
	return ok
}

// IsAudit is audit
func IsAudit(db *gorm.DB) bool {
	field := db.Statement.Schema.LookUpField(FieldNameTriggerID)
	return field != nil
}

// IsAuditableField is auditable field
func IsAuditableField(field *schema.Field) bool {
	value, ok := field.Tag.Lookup(AuditableTag)
	return ok && value == AuditableTagValueField
}

// IsAuditableLogsField is auditable logs field
func IsAuditableLogsField(field *schema.Field) bool {
	value, ok := field.Tag.Lookup(AuditableTag)
	return ok && value == AuditableTagValueLogs
}
