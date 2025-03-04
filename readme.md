# Tg hotels bot

## Функционал

- БД **облачный сервис** [MongoDB Atlas](https://www.mongodb.com/atlas/database).
- [RapidAPI](https://rapidapi.com/apidojo/api/hotels4/)
- [Unsplash](https://unsplash.com)

...

## Notes

### Подготовка проекта

- Удаляем кэш старой сборки

```bash
go clean -cache
```

- Соберет бинарник в .bin/bot

```bash
make build
```

- Соберет и запустит бота

```bash
make run
```

### Линтер

```bash
golangci-lint run
```

## pre-commit

1. Установка pre-commit

```bash
pip install pre-commit
```

2. Установка хуков

```bash
pre-commit install
```

### Запуск

```bash
pre-commit run -a
```
