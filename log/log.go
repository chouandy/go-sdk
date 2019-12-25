package log

import (
	"errors"
	"fmt"
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
func Init(config Config) (err error) {
	// Print log
	TextLog().WithFields(config.LogrusFields()).Info("init log")

	// Set formatter
	switch config.Format {
	case "text":
		Log.SetFormatter(&logrus.TextFormatter{
			CallerPrettyfier: callerPrettyfier,
		})
	case "json":
		Log.SetFormatter(&logrus.JSONFormatter{
			CallerPrettyfier: callerPrettyfier,
			FieldMap:         fieldMap,
		})
	default:
		err = errors.New("unknown formatter")
		return
	}

	// Set level
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		return err
	}
	Log.SetLevel(level)

	// Set reporter caller
	Log.SetReportCaller(true)

	// Add hooks
	Log.AddHook(hooksex.NewSkipCallerHook())

	return nil
}

func callerPrettyfier(f *runtime.Frame) (function string, file string) {
	function = fmt.Sprintf("%s()", f.Function)
	file = fmt.Sprintf("%s:%d", path.Base(f.File), f.Line)
	return
}
