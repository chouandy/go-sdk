package mailer

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"

	osex "github.com/chouandy/go-sdk/os"
)

// Config config
type Config struct {
	SMTPSettings SMTPSettings
	Options      Options
	URLOptions   URLOptions
}

// SMTPSettings smtp settings
type SMTPSettings struct {
	Address  string
	Port     int
	Username string
	Password string
}

// Options options struct
type Options struct {
	From string
}

// URLOptions url options struct
type URLOptions struct {
	Protocol string
	Host     string
}

// NewConfig new config
func NewConfig() (*Config, error) {
	// New config
	config := Config{
		SMTPSettings: SMTPSettings{
			Address:  os.Getenv("MAILER_SMTP_ADDRESS"),
			Port:     osex.GetenvParseInt("MAILER_SMTP_PORT"),
			Username: os.Getenv("MAILER_SMTP_USERNAME"),
			Password: os.Getenv("MAILER_SMTP_PASSWORD"),
		},
		Options: Options{
			From: os.Getenv("MAILER_FROM"),
		},
		URLOptions: URLOptions{
			Protocol: os.Getenv("MAILER_PROTOCOL"),
			Host:     os.Getenv("MAILER_HOST"),
		},
	}
	// Validate driver
	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &config, nil
}

// Validate validate
func (c *Config) Validate() error {
	if len(c.SMTPSettings.Address) == 0 {
		return errors.New("smtp address can't be blank")
	}
	if c.SMTPSettings.Port == 0 {
		return errors.New("smtp port can't be blank")
	}
	if len(c.SMTPSettings.Username) == 0 {
		return errors.New("smtp username can't be blank")
	}
	if len(c.SMTPSettings.Password) == 0 {
		return errors.New("smtp password can't be blank")
	}
	if len(c.Options.From) == 0 {
		return errors.New("from can't be blank")
	}
	if len(c.URLOptions.Protocol) == 0 {
		return errors.New("url protocal can't be blank")
	}
	if len(c.URLOptions.Host) == 0 {
		return errors.New("url host can't be blank")
	}

	return nil
}

// URL url
func (c *Config) URL() string {
	return c.URLOptions.Protocol + "://" + c.URLOptions.Host
}

// LogrusFields logrus fields
func (c *Config) LogrusFields() logrus.Fields {
	return logrus.Fields{
		"smtp_settings": logrus.Fields{
			"address": c.SMTPSettings.Address,
			"port":    c.SMTPSettings.Port,
		},
		"options": logrus.Fields{
			"from": c.Options.From,
		},
		"url_options": logrus.Fields{
			"protocol": c.URLOptions.Protocol,
			"host":     c.URLOptions.Host,
		},
	}
}
