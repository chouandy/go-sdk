package config

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v2"

	logex "github.com/chouandy/go-sdk/log"
	osex "github.com/chouandy/go-sdk/os"
)

// Stage stage
var Stage = osex.Getenv("STAGE", "local")

var configDir = "config"

// Load load config
func Load(config interface{}, debug bool) error {
	// Get config reflect value
	v := reflect.Indirect(reflect.ValueOf(config))

	// Interate config all fields
	for i := 0; i < v.Type().NumField(); i++ {
		// Get config file name with field name to snake case
		name := strcase.ToSnake(v.Type().Field(i).Name)

		// New filename
		filename := fmt.Sprintf("%s/%s/%s.yml", configDir, Stage, name)

		// Read config file
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}

		// Unmarshal config file to field struct
		if err := yaml.Unmarshal(data, v.Field(i).Addr().Interface()); err != nil {
			return err
		}

		if !debug {
			// Print log
			logex.Log.Infof("load %s", filename)
		} else {
			// Print log
			logex.Log.WithFields(logex.ToLogrusFields(v.Field(i).Interface())).Infof("load %s", filename)
		}
	}

	return nil
}

// SetConfigDir set config dir
func SetConfigDir(dir string) {
	configDir = dir
}
