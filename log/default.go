package log

import (
	"github.com/sirupsen/logrus"

	hooksex "github.com/chouandy/go-sdk/log/hooks"
)

// TextLog default log
var textLog *logrus.Logger

// TextLog return default text log
func TextLog() *logrus.Logger {
	// Check text log is new or not
	if textLog == nil {
		// New log
		textLog = logrus.New()
		// Set formatter
		textLog.SetFormatter(&logrus.TextFormatter{
			CallerPrettyfier: callerPrettyfier,
		})
		// Set level
		textLog.SetLevel(logrus.InfoLevel)
		// Set reporter caller
		textLog.SetReportCaller(true)
		// Add hooks
		textLog.AddHook(hooksex.NewSkipCallerHook())
	}

	return textLog
}
