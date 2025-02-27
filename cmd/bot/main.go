package main

import (
	"context"
	"log"
	"tg-hotels-bot/internal/config"
	"tg-hotels-bot/internal/database"
	"tg-hotels-bot/internal/rapidapi/rapidapi_requests"
	"tg-hotels-bot/internal/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Printf("Ошибка загрузки конфигурации: %v", err)
		return
	}

	// Подключаем MongoDB
	// никуда не передаем, а просто проверяем подключение // FIXME?
	mongoClient, err := database.GetMongoClient(cfg)
	if err != nil {
		log.Printf("Ошибка подключения к MongoDB: %v", err)
		return
	}
	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			log.Println("Ошибка при закрытии соединения с MongoDB: ", err)
			return
		}
	}()

	// Загружаем мета-данные RapidAPI
	log.Println("Запрашиваем мета-данные с RapidAPI...")
	if err := rapidapi_requests.FetchMetaData(); err != nil {
		log.Println("Ошибка при загрузке мета-данных: ", err)
		return
	}

	botApi, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Println("Ошибка создания бота: ", err)
		return
	}

	bot := telegram.NewBot(botApi)

	bot.Start(cfg)
}
