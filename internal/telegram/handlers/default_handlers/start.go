package default_handlers

import (
	reply_keyboards "tg-hotels-bot/internal/keyboards/reply"
	"tg-hotels-bot/internal/photos"
	"tg-hotels-bot/internal/states"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleStart(bot *tgbotapi.BotAPI, message *tgbotapi.Message, stateManager *states.StateManager) {
	photo := tgbotapi.NewPhoto(message.Chat.ID, tgbotapi.FileURL(photos.Photos["home_menu"]))
	photo.Caption = "<b>Выберите действие</b>"
	photo.ParseMode = "HTML"
	photo.ReplyMarkup = reply_keyboards.HomeMenuKeyboard()
	bot.Send(photo)

	stateManager.ClearState(message.Chat.ID)
}

func HandleMainMenu(bot *tgbotapi.BotAPI, message *tgbotapi.Message, stateManager *states.StateManager) {
	HandleStart(bot, message, stateManager)
}
