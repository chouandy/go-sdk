package log

import "github.com/sirupsen/logrus"

// ToLogrusFields to logrus fields
func ToLogrusFields(v interface{}) logrus.Fields {
	// New fields
	fields := make(logrus.Fields)

	// v to json
	if data, err := json.Marshal(&v); err == nil {
		// json to fields
		json.Unmarshal(data, &fields)
	}

	return fields
}
