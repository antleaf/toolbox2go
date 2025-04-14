package toolbox2go

import "os"

import (
	"gopkg.in/yaml.v3"
)

func NewConfigurationFromYamlFile[T any](config T, configFilePath string) error {
	configData, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return err
	}
	return err
}
