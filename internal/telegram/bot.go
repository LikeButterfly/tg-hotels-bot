package telegram

import (
	"log"

	"tg-hotels-bot/internal/config"
	"tg-hotels-bot/internal/states"
	"tg-hotels-bot/internal/telegram/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	BotAPI       *tgbotapi.BotAPI
	StateManager *states.StateManager
	UserData     map[int64]map[string]string // Хранение данных пользователей
}

// NewBot создает новый экземпляр бота
func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{
		BotAPI:       bot,
		StateManager: states.NewStateManager(),
		UserData:     make(map[int64]map[string]string),
	}
}

// Start запускает бота
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

	// Передаем stateManager и userData в обработчик команд
	handlers.HandleCommands(b.BotAPI, updates, b.StateManager, b.UserData)
}
