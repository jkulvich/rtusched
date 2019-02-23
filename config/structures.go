package config

import "rtusched/logger"

type Config struct {
	Site string `yaml:"site"`
	Log logger.Config `yaml:"log"`
}