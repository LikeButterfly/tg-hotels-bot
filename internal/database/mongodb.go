package database

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"tg-hotels-bot/internal/config"
)

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
	clientErr      error
)

func GetMongoClient(cfg *config.Config) (*mongo.Client, error) {
	clientOnce.Do(func() {
		clientInstance, clientErr = connect(cfg)
	})
	return clientInstance, clientErr
}

func connect(cfg *config.Config) (*mongo.Client, error) {
	// if cfg.MongoDBUser == "" || cfg.MongoDBPass == "" {
	// 	return nil, errors.New("отсутствуют учетные данные для MongoDB")
	// }

	uri := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
		cfg.MongoDBUser,
		cfg.MongoDBPass,
		cfg.MongoDBHost,
	)

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к MongoDB: %w", err)
	}

	// Проверка подключения
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("MongoDB не отвечает: %w", err)
	}

	log.Println("Успешное подключение к MongoDB")
	return client, nil
}
