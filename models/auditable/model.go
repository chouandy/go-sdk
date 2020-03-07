package auditable

import modelsex "github.com/chouandy/go-sdk/models"

// Interface is used to get metadata from your models.
type Interface interface {
	IsAuditable() bool
}

// Model auditable model struct
type Model struct {
	OriginalEntity interface{} `gorm:"-"`
	TriggerID      *uint64     `gorm:"-"`
}

// IsAuditable is auditable
func (Model) IsAuditable() bool {
	return true
}

// SetTriggerID set trigger id
func (m *Model) SetTriggerID(id uint64) {
	m.TriggerID = modelsex.UInt64(id)
}
