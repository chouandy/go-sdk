package log

import "github.com/sirupsen/logrus"

// Config config struct
type Config struct {
	Format string `json:"format" yaml:"format"`
	Level  string `json:"level" yaml:"level"`
}

// LogrusFields logrus fields
func (c *Config) LogrusFields() logrus.Fields {
	return logrus.Fields{
		"format": c.Format,
		"level":  c.Level,
	}
}
