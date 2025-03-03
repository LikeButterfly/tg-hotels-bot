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