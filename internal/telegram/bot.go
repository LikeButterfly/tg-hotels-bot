package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tg-hotels-bot/internal/config"
	"tg-hotels-bot/internal/models"
	"tg-hotels-bot/internal/states"
	"tg-hotels-bot/internal/telegram/handlers"
)

// Создает новый экземпляр бота
func NewBot(bot *tgbotapi.BotAPI) *models.Bot {
	return &models.Bot{
		BotAPI:       bot,
		StateManager: states.NewStateManager(),
		UsersData:    make(map[int64]*models.UserData),
	}
}

// Запускает бота
func Start(cfg *config.Config, bot *models.Bot) {
	bot.BotAPI.Debug = false
	log.Printf("Авторизован как %s", bot.BotAPI.Self.UserName)

	SetDefaultCommands(cfg, bot)

	// Настраиваем поллер (длинные запросы)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 20

	updates := bot.BotAPI.GetUpdatesChan(u)

	handlers.HandleCommands(cfg, bot.BotAPI, updates, bot.StateManager, bot.UsersData)
}
