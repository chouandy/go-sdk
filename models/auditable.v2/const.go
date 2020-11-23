package auditable

// Tags
const (
	AuditableTag           = "auditable"
	AuditableTagValueField = "field"
	AuditableTagValueLogs  = "logs"
)

// FieldNames
const (
	FieldNameOriginalEntity = "OriginalEntity"
	FieldNameTriggerID      = "TriggerID"
	FieldNameAuditableID    = "AuditableID"
	FieldNameAction         = "Action"
	FieldNameChanges        = "Changes"
)

// Actions
const (
	ActionCreate uint32 = 0
	ActionUpdate uint32 = 1
	ActionDelete uint32 = 2
)
