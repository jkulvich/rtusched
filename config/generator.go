package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"rtusched/logger"
)

//GenerateDefault - Генериурет стандартный файл конфигурации по указаному пути и возвращает содержимое
func GenerateDefault(filename string) (*Config, error) {
	conf := Config{
		Site: "https://www.mirea.ru/education/schedule-main/schedule/",
		Log: logger.Config{
			Level: "warn",
			Format: "text",
			CallerInfo: false,
		},
	}
	data, err := yaml.Marshal(conf)
	if err != nil {
		return nil, err
	}
	if err := ioutil.WriteFile(filename, data, 0777); err != nil {
		return nil, err
	}
	return &conf, nil
}
