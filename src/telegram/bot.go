package telegram

import (
	"log"
	"tg-hotels-bot/src/config"
	"tg-hotels-bot/src/telegram/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	BotAPI *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{BotAPI: bot}
}

func (b *Bot) Start(cfg *config.Config) {
	b.BotAPI.Debug = true
	log.Printf("Авторизован как %s", b.BotAPI.Self.UserName)

	b.SetDefaultCommands(cfg)

	// Настраиваем поллер (длинные запросы)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 20

	updates, err := b.BotAPI.GetUpdatesChan(u)

	if err != nil {
		log.Fatal("Ошибка получения обновлений:", err)
	}

	handlers.HandleCommands(b.BotAPI, updates)

}
