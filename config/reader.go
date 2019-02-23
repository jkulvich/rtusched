package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//ReadFile - Читает конфигурацию из файла
func ReadFile(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var conf = &Config{}
	if err := yaml.Unmarshal(data, conf); err != nil {
		return nil, err
	}
	return conf, nil
}