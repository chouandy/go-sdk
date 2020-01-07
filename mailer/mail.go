package mailer

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"gopkg.in/gomail.v2"

	logex "github.com/chouandy/go-sdk/log"
)

// Mail mail struct
type Mail struct {
	To           string                 `json:"to"`
	Subject      string                 `json:"subject"`
	TemplatePath string                 `json:"template_path"`
	TemplateName string                 `json:"template_name"`
	Variables    map[string]interface{} `json:"variables"`
}

// TemplateFilePath return template file relative path
func (m *Mail) TemplateFilePath() string {
	return fmt.Sprintf("%s/%s", m.TemplatePath, m.TemplateName)
}

// GetBody get body
func (m *Mail) GetBody() (string, error) {
	// Read mail body template
	body, err := Load(m.TemplateFilePath())
	if err != nil {
		logex.Log.Error(err)
		return "", err
	}

	// Parse email body
	tmpl, err := template.New("email").Parse(string(body))
	if err != nil {
		logex.Log.Error(err)
		return "", err
	}
	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, m.Variables); err != nil {
		logex.Log.Error(err)
		return "", err
	}

	return buffer.String(), nil
}

// ContentType get content type by template name ext
func (m *Mail) ContentType() string {
	if strings.HasSuffix(m.TemplateName, "html") {
		return "text/html"
	}

	return "text/plain"
}

// Send send
func (m *Mail) Send() error {
	// New mail message
	msg := gomail.NewMessage()
	// Set from
	msg.SetHeader("From", config.Options.From)
	// Set mail to
	msg.SetHeader("To", m.To)
	// Set mail subject
	msg.SetHeader("Subject", m.Subject)

	// Get mail body
	body, err := m.GetBody()
	if err != nil {
		logex.Log.Error(err)
		return err
	}
	// Set email body
	msg.SetBody(m.ContentType(), body)

	// New dialer
	d := gomail.NewDialer(
		config.SMTPSettings.Address,
		config.SMTPSettings.Port,
		config.SMTPSettings.Username,
		config.SMTPSettings.Password,
	)

	// Send email
	if err := d.DialAndSend(msg); err != nil {
		logex.Log.Error(err)
		return err
	}

	return nil
}
