package default_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	reply_keyboards "tg-hotels-bot/internal/keyboards/reply"
	"tg-hotels-bot/internal/states"
	"tg-hotels-bot/internal/unsplash"
)

func HandleStart(bot *tgbotapi.BotAPI, message *tgbotapi.Message, stateManager *states.StateManager) {
	photo := tgbotapi.NewPhoto(message.Chat.ID, tgbotapi.FileURL(unsplash.GetPhotoByCommand("home_menu")))
	photo.Caption = "<b>Выберите действие</b>"
	photo.ParseMode = "HTML"
	photo.ReplyMarkup = reply_keyboards.HomeMenuKeyboard()
	if _, err := bot.Send(photo); err != nil {
		// TODO обработать ошибку
		return
	}
	stateManager.ClearState(message.Chat.ID)
}

func HandleMainMenu(bot *tgbotapi.BotAPI, message *tgbotapi.Message, stateManager *states.StateManager) {
	HandleStart(bot, message, stateManager)
}
