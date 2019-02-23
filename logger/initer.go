package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

//Init - настраивает новый логер
func Init(conf Config) (*logrus.Logger, error) {
	logger := logrus.New()

	// Уровень логирования
	level, err := logrus.ParseLevel(conf.Level)
	if err != nil {
		return nil, err
	}
	logger.SetLevel(level)

	// Формат логов
	switch conf.Format {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		logger.SetFormatter(&logrus.TextFormatter{})
	default:
		return nil, fmt.Errorf("unknown log formatter '%s', try json or text", conf.Format)
	}

	// Дополнительная информация о вызвавшей функции
	logger.SetReportCaller(conf.CallerInfo)

	return logger, nil
}
