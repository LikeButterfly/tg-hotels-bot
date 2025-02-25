package handlers

import (
	"tg-hotels-bot/src/telegram/handlers/default_handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleCommands(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Text {
		case "/start":
			default_handlers.HandleStart(bot, update.Message)
		case "Главное меню":
			default_handlers.HandleMainMenu(bot, update.Message)
		case "/help":
			default_handlers.HandleHelp(bot, update.Message)
		case "Справка":
			default_handlers.HandleHelpText(bot, update.Message)
		}
	}
}
