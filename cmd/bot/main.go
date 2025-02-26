package main

import (
	"context"
	"log"

	"tg-hotels-bot/internal/config"
	"tg-hotels-bot/internal/database"
	"tg-hotels-bot/internal/rapidapi/rapidapi_requests"
	"tg-hotels-bot/internal/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	cfg := config.LoadConfig()

	// Подключаем MongoDB
	client := database.GetMongoClient()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal("Ошибка при закрытии соединения с MongoDB:", err)
		}
	}()

	// Загружаем мета-данные RapidAPI
	log.Println("Запрашиваем мета-данные с RapidAPI...")
	if err := rapidapi_requests.FetchMetaData(); err != nil {
		log.Fatal("Ошибка при загрузке мета-данных:", err)
	}

	botApi, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Fatal("Ошибка создания бота:", err)
	}

	bot := telegram.NewBot(botApi)

	bot.Start(cfg)
}
