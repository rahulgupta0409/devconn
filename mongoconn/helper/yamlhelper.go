package helper

import (
	"fmt"
	"log"
	"os"

	"github.com/rahulgupta0409/devconn/models"
	"gopkg.in/yaml.v3"
)

func Yamlfile(filename string) string {
	return filename
}

func Readyaml(filepath string) models.Collection {
	var conn models.Collection

	// Open YAML file
	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()

	// Decode YAML file to struct
	if file != nil {
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&conn); err != nil {
			log.Println(err.Error())
		}
	}
	fmt.Println(conn)
	return conn
}
