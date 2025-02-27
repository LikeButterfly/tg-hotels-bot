package work_with_messages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendWaitingMessage(bot *tgbotapi.BotAPI, chatID int64) (int, int) {
	textMsg := tgbotapi.NewMessage(chatID, "<i>Выполняю поиск...</i>")
	textMsg.ParseMode = "HTML"
	sentText, _ := bot.Send(textMsg)

	stickerMsg := tgbotapi.NewSticker(chatID, tgbotapi.FileID("CAACAgIAAxkBAAEGPdBjXOSiezHXCIYn3k_DU2Nx_khoVwAC6BUAAiMlyUtQqGgG1fAXAAEqBA"))
	sentSticker, _ := bot.Send(stickerMsg)

	return sentText.MessageID, sentSticker.MessageID
}

func DeleteWaitingMessages(bot *tgbotapi.BotAPI, chatID int64, textMsgID, stickerMsgID int) {
	DeleteMessage(bot, chatID, textMsgID)
	DeleteMessage(bot, chatID, stickerMsgID)
}
