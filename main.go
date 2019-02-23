package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"rtusched/config"
	"rtusched/logger"
)

var BuildVersion = "unmarked"
var BuildType = "debug"

func main() {
	const configDefaultFilename = "rtusched.yaml"
	fmt.Printf("RTU Scheduled, Version: %s - %s\n", BuildVersion, BuildType)

	// Парсинг флагов
	flagConfigPath := flag.String("config", configDefaultFilename, "Sets config file path")
	flag.Parse()

	// Получение глобальной конфигурации
	// Пытаемся получить конфиг из файла. Генерируем стандартный в случае неудачи
	conf, err := config.ReadFile(*flagConfigPath)
	if os.IsNotExist(err) {
		conf, err = config.GenerateDefault(*flagConfigPath)
		if err != nil {
			log.Fatalf("can't generate default config, check rights: %s", err)
		}
		log.Println("new default config was generated and used")
	}
	if err != nil {
		log.Fatalf("can't read config file: %s", err)
	}

	// Получение логера
	logx, err := logger.Init(conf.Log)
	if err != nil {
		log.Fatalf("can't init logger: %s", err)
	}

	logx.Warn("OK!")
}
