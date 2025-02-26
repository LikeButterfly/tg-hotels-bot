package database

import (
	"context"
	"fmt"
	"log"
	"sync"
	"tg-hotels-bot/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
)

func GetMongoClient() *mongo.Client { // TODO лучше возвращать еще и error
	clientOnce.Do(func() {
		cfg := config.LoadConfig()
		uri := fmt.Sprintf(
			"mongodb+srv://%s:%s@gohotelsbot.ctzeo.mongodb.net/?retryWrites=true&w=majority", // FIXME - gohotelsbot в env
			// "mongodb+srv://%s:%s@gohotelsbot.ctzeo.mongodb.net/?retryWrites=true&w=majority&tls=true&appName=GoHotelsBot", // FIXME - GoHotelsBot в env
			cfg.MongoDBUser,
			cfg.MongoDBPass,
		)

		clientOptions := options.Client().ApplyURI(uri)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatalf("Ошибка подключения к MongoDB: %v", err)
		}

		// Проверка подключения
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatalf("MongoDB не отвечает: %v", err)
		}

		clientInstance = client
		log.Println("Успешное подключение к MongoDB")
	})

	return clientInstance
}
