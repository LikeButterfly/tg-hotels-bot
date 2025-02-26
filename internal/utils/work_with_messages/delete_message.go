package work_with_messages

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func DeleteMessage(bot *tgbotapi.BotAPI, chatID int64, messageID int) {
	deleteMsg := tgbotapi.NewDeleteMessage(chatID, messageID)
	_, err := bot.Send(deleteMsg)
	if err != nil {
		log.Printf("Ошибка при удалении сообщения %d в чате %d: %v", messageID, chatID, err)
	}
}
