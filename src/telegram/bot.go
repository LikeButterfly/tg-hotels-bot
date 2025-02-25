package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start() {
	b.bot.Debug = true
	log.Printf("Авторизован как %s", b.bot.Self.UserName)

	// Настраиваем поллер (длинные запросы)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 20

	updates, err := b.bot.GetUpdatesChan(u)

	if err != nil {
		log.Fatal("Ошибка получения обновлений:", err)
	}

	if err := b.HandleCommands(updates); err != nil {
		log.Fatal("Ошибка обработки команд:", err)
	}
}
