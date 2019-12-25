package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// NotLoggedPaths not logged paths
var NotLoggedPaths map[string]struct{}

// GetLogLevel get log level by status code
func GetLogLevel(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "INFO"
	case code >= 300 && code < 400:
		return "INFO"
	case code >= 400 && code < 500:
		return "ERROR"
	default:
		return "FATAL"
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// LoggerWithNotLogged logger with not logged
func LoggerWithNotLogged(logger *logrus.Logger, paths ...string) gin.HandlerFunc {
	// Set not logged
	if length := len(paths); length > 0 {
		NotLoggedPaths = make(map[string]struct{}, length)
		for _, path := range paths {
			NotLoggedPaths[path] = struct{}{}
		}
	}

	return LoggerWithLogrus(logger)
}

// LoggerWithLogrus logger with writer
func LoggerWithLogrus(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Replace context writer
		writer := &bodyLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = writer

		// Process request
		c.Next()

		// Not logged
		if _, ok := NotLoggedPaths[c.Request.URL.Path]; ok {
			return
		}

		// Get status code
		status := c.Writer.Status()

		// New logger fields
		fields := logrus.Fields{
			"status":    status,
			"method":    c.Request.Method,
			"path":      c.Request.URL.Path,
			"latency":   fmt.Sprintf("%v", time.Now().UTC().Sub(start)),
			"client_ip": c.ClientIP(),
		}
		// Set request id
		if requestID, exists := c.Get("request_id"); exists {
			fields["request_id"] = requestID.(string)
		}
		// Set query string parameters
		if len(c.Request.URL.RawQuery) > 0 {
			fields["query_string_parameters"] = c.Request.URL.Query()
		}
		// Set path parameters
		if len(c.Params) > 0 {
			params := map[string]string{}
			for _, param := range c.Params {
				params[param.Key] = param.Value
			}
			fields["path_parameters"] = params
		}
		// Set request body
		if body, err := c.GetRawData(); err == nil && len(body) > 0 {
			fields["body"] = string(body)
		}
		// Set error
		if status >= http.StatusBadRequest {
			fields["error"] = json.RawMessage(writer.body.Bytes())
		}
		// Set location
		if status == http.StatusFound {
			fields["location"] = c.Writer.Header().Get("Location")
		}

		// Set logger fields
		entry := logger.WithFields(fields)

		// Log by status code
		if status >= 200 && status < 400 {
			entry.Info()
		} else if status >= 400 && status < 500 {
			entry.Error()
		} else {
			entry.Fatal()
		}
	}
}
