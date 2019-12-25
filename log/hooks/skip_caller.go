package hooks

import "github.com/sirupsen/logrus"

// SkipCallerHook caller hook struct
type SkipCallerHook struct{}

// NewSkipCallerHook new caller hook
func NewSkipCallerHook() *SkipCallerHook {
	return new(SkipCallerHook)
}

// Fire fire
func (h *SkipCallerHook) Fire(entry *logrus.Entry) error {
	entry.Caller = nil
	return nil
}

// Levels return applied levels
func (h *SkipCallerHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.InfoLevel,
	}
}
