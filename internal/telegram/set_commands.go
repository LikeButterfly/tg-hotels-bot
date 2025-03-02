package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tg-hotels-bot/internal/config"
	"tg-hotels-bot/internal/models"
)

func SetDefaultCommands(cfg *config.Config, bot *models.Bot) {
	var commands []tgbotapi.BotCommand
	for _, cmd := range config.DefaultCommands {
		commands = append(commands, tgbotapi.BotCommand{
			Command:     cmd.Command,
			Description: cmd.Desc,
		})
	}

	setCmdCfg := tgbotapi.NewSetMyCommands(commands...)
	_, err := bot.BotAPI.Request(setCmdCfg)
	if err != nil {
		log.Println("Ошибка установки команд:", err)
		return
	}

	log.Println("Команды успешно установлены")
}
