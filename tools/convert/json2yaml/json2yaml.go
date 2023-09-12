package json2yaml

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Json2Yaml interface {
	Convert(filename string) (string, error)
}

type Json2YamlImpl struct {
}

const fileExt = ".yaml"

func (obj *Json2YamlImpl) Convert(filename string) (string, error) {
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	var jsonObj interface{}
	err = json.Unmarshal(jsonData, &jsonObj)
	if err != nil {
		return "", nil
	}

	yamlData, err := yaml.Marshal(jsonObj)
	if err != nil {
		log.Fatal(err)
	}

	fileName := filepath.Base(filename)
	outputFileName := strings.TrimSuffix(fileName, filepath.Ext(fileName)) + fileExt
	err = os.WriteFile(outputFileName, yamlData, 0644)
	if err != nil {
		return "", err
	}
	return outputFileName, nil
}
