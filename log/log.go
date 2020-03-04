package log

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"

	hooksex "github.com/chouandy/go-sdk/log/hooks"
)

// Log is logrus logger
var Log = logrus.New()

// Custom fields key
var fieldMap = logrus.FieldMap{
	logrus.FieldKeyTime:  "@timestamp",
	logrus.FieldKeyLevel: "@level",
	logrus.FieldKeyMsg:   "@message",
	logrus.FieldKeyFunc:  "@func",
	logrus.FieldKeyFile:  "@file",
}

// Init init
func Init() {
	// Set formatter
	SetFormatter()
	// Set level
	SetLevel()
	// Set reporter caller
	Log.SetReportCaller(true)
	// Add hooks
	Log.AddHook(hooksex.NewSkipCallerHook())
}

// SetFormatter set formatter
func SetFormatter() {
	// Check env
	if os.Getenv("LOG_FORMATTER") == "json" {
		Log.SetFormatter(&logrus.JSONFormatter{
			CallerPrettyfier: callerPrettyfier,
			FieldMap:         fieldMap,
		})
	}

	Log.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: callerPrettyfier,
	})
}

// SetLevel set level
func SetLevel() {
	// Check env
	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		level = logrus.DebugLevel
	}

	// Set level
	Log.SetLevel(level)
}

func callerPrettyfier(f *runtime.Frame) (function string, file string) {
	function = fmt.Sprintf("%s()", f.Function)
	file = fmt.Sprintf("%s:%d", path.Base(f.File), f.Line)
	return
}
