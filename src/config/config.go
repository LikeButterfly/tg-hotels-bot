package config

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken        string
	MongoDBUser     string
	MongoDBPass     string
	TimeOffset      time.Duration
	DefaultCommands map[string]string
}

var (
	cfg  *Config
	once sync.Once // ленивая загрузка, то есть один раз
)

func LoadConfig() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Ошибка загрузки .env файла")
			// FIXME ?
		}

		cfg = &Config{
			BotToken:    getEnv("BOT_TOKEN", ""),
			MongoDBUser: getEnv("MONGO_DB_USERNAME", ""),
			MongoDBPass: getEnv("MONGO_DB_PASSWORD", ""),
			TimeOffset:  3 * time.Hour,
			DefaultCommands: map[string]string{
				"start":    "Запуск",
				"help":     "Помощь",
				"lowprice": "Топ самых дешёвых отелей в городе",
				// "highprice": "Топ самых дорогих отелей в городе",                                // no support
				// "bestdeal":  "Топ отелей, наиболее подходящих по цене и расположению от центра", // no support
				// "history": "История", // no support
				// "favorites": "Избранное", // no support
			},
		}
	})
	return cfg
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value // FIXME ?
	}
	return defaultValue
}
