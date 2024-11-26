package database

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

func LoadQuery(fileName string) string {
	path := filepath.Join("internal", "database", "sql", fileName)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to load query %s: %v", fileName, err)
	}

	return string(content)
}