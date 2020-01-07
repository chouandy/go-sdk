package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	logex "github.com/chouandy/go-sdk/log"
)

var sqlDir = "sql"

var sqls = map[string]string{}

// LoadSQLFiles load sql files
func LoadSQLFiles() error {
	// Check sql dir is exist or not
	if _, err := os.Stat(sqlDir); os.IsNotExist(err) {
		return nil
	}

	// Read sql files
	return readSQLFiles(sqlDir)
}

func readSQLFiles(dir string) error {
	// Get sql dir files
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
			// Read sub dir sql files
			readSQLFiles(filename)
		} else {
			// Check ext
			if ext := filepath.Ext(file.Name()); ext != ".sql" {
				continue
			}

			// Read file content
			data, err := ioutil.ReadFile(filename)
			if err != nil {
				return err
			}

			// Set to sqls map
			sqls[filename] = string(data)

			logex.TextLog().Infof("load %s", filename)
		}
	}

	return nil
}

// Load return sql content
func Load(filePath string) string {
	return sqls[fmt.Sprintf("%s/%s.sql", sqlDir, filePath)]
}

// SetSQLDir set sql dir
func SetSQLDir(dir string) {
	sqlDir = dir
}
