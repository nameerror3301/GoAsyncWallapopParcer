# GoAsyncWallapopParcer

## Описание

Микросервис для сбора данных с сайта Wallapop.es

Логика работы проста, параллельно из 5 категорий собираются данные и отправляются в другой сервис каждую минуту. 

Обмен данными будет происходить по средствами JSON объектов. Каждая категория это отдельный объект.

## Структура проекта

```
.
├── app
│   └── cmd
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   └── run.go
│   └── models
│       └── wallapop.go // Бизнес логика
├── Readme.md
```

## Использование

На данном этапе разработки, пока сервис не работает в контейнере, вам потребуется иметь на своем устройстве **Golang версии 1.19**. 

Требуется перейти в папку **/app/cmd/** и запустить файл **main.go**. После чего в консоли вы увидите результат.
