package mailer

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	logex "github.com/chouandy/go-sdk/log"
)

var templateDir = "static/mailer"

var templates = map[string]string{}

// LoadTemplateFiles load template files
func LoadTemplateFiles() error {
	// Check template dir is exist or not
	if _, err := os.Stat(templateDir); os.IsNotExist(err) {
		return nil
	}

	// Read template files
	return readTemplateFiles(templateDir)
}

func readTemplateFiles(dir string) error {
	// Get template dir files
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	// Interate files
	for _, file := range files {
		// New file name
		filename := fmt.Sprintf("%s/%s", dir, file.Name())
		// Check is dir or not
		if file.IsDir() {
			// Read sub dir template files
			readTemplateFiles(filename)
		} else {
			// Read file content
			data, err := ioutil.ReadFile(filename)
			if err != nil {
				return err
			}

			// Set to templates map
			templates[filename] = string(data)

			logex.Log.Infof("load %s", filename)
		}
	}

	return nil
}

// Load return template content
func Load(filePath string) (string, error) {
	content, ok := templates[fmt.Sprintf("%s/%s", templateDir, filePath)]
	if !ok {
		return "", errors.New("template not found")
	}

	return content, nil
}

// SetTemplateDir set template dir
func SetTemplateDir(dir string) {
	templateDir = dir
}
