package config

import (
	"errors"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken    string
	MongoDBUser string
	MongoDBPass string
	MongoDBHost string
}

var (
	cfg       *Config
	once      sync.Once // ленивая загрузка, то есть один раз
	configErr error
)

func LoadConfig() (*Config, error) {
	once.Do(func() {
		cfg, configErr = load()
	})
	return cfg, configErr
}

func load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("ошибка загрузки .env файла: " + err.Error())
	}

	config := &Config{
		BotToken:    getEnv("BOT_TOKEN"),
		MongoDBUser: getEnv("MONGO_DB_USERNAME"),
		MongoDBPass: getEnv("MONGO_DB_PASSWORD"),
		MongoDBHost: getEnv("MONGO_DB_HOST"),
	}

	if config.BotToken == "" {
		return nil, errors.New("отсутствует BOT_TOKEN в переменных окружения")
	}

	if config.MongoDBUser == "" || config.MongoDBPass == "" || config.MongoDBHost == "" {
		return nil, errors.New("отсутствуют учетные данные MongoDB")
	}

	return config, nil
}

func getEnv(key string) string {
	return os.Getenv(key)
}
