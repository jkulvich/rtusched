# RTU Scheduled
Микросервис предоставляющий более удобное RESTfull API для получения данных расписания  
Расписание РТУ по умолчанию находится на сайте: [mirea.ru/education/schedule-main/schedule](https://www.mirea.ru/education/schedule-main/schedule/)  

Сервис парсит все xlsx расписания с сайта и формирует нормальное json представление для каждого из них.
Далее может быть развёрнут внутренний сервер предоставляющий REST доступ к расписанию.  

# Сборка проекта
0. Клонируйте репозиторий и перейдите в него
1. [Установите GoLang](https://golang.org/dl/)
2. Выполните `go build` или `go build -ldflags "-X main.BuildVersion={version} -X main.BuildType=release"`

# Запуск
- Сервис готов к работе сразу после сборки  
- Базовая конфигурация будет создана автоматически (предоставьте права на изменения в текущей директории)
- Имя станартного файла конфигурации: `rtusched.yaml`  
- Указать пользовательский файл конфигурации можно флагом `--config filename`

## Конфигурация
YAML
```yaml
# Страница с расписанием РТУ
site: https://www.mirea.ru/education/schedule-main/schedule/

# Конфигурция логировщика
log:
  level: warn          #< Уровень логирования: panic, fatal, error, warn, info, debug, trace
  format: text         #< Формат вывода: text, json
  callerInfo: false    #< Дополнительная информация о методе и номере строки
```

# API
> Скоро будет...

# Q&A
> Скоро будет...

