package auditable

import "github.com/jinzhu/gorm"

// IsAuditableModel is auditable model
func IsAuditableModel(scope *gorm.Scope) bool {
	_, ok := scope.Value.(Interface)
	return ok
}

// IsAudit is audit
func IsAudit(scope *gorm.Scope) bool {
	f, ok := scope.FieldByName(FieldNameTriggerID)
	return ok && !f.Field.IsNil()
}

// IsAuditableField is auditable field
func IsAuditableField(f *gorm.Field) bool {
	value, ok := f.Tag.Lookup(AuditableTag)
	return ok && value == AuditableTagValueField
}

// IsAuditableLogsField is auditable logs field
func IsAuditableLogsField(f *gorm.Field) bool {
	value, ok := f.Tag.Lookup(AuditableTag)
	return ok && value == AuditableTagValueLogs
}
