package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func LoadYamlConfig(filePath string, obj interface{}) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal([]byte(file), obj); err != nil {
		return err
	}
	return nil
}
