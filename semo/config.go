package semo

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Description string  `yaml:"description"`
	Fields      []Field `yaml:"fields"`
}

func NewConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	dec := yaml.NewDecoder(file)

	var config Config
	err = dec.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
