package default_handlers

import (
	reply_keyboards "tg-hotels-bot/src/keyboards/reply"
	"tg-hotels-bot/src/photos"
	"tg-hotels-bot/src/states"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleStart(bot *tgbotapi.BotAPI, message *tgbotapi.Message, stateManager *states.StateManager) {
	photo := tgbotapi.NewPhotoShare(message.Chat.ID, photos.Photos["home_menu"])
	photo.Caption = "<b>Выберите действие</b>"
	photo.ParseMode = "HTML"
	photo.ReplyMarkup = reply_keyboards.HomeMenuKeyboard()
	bot.Send(photo)

	stateManager.ClearState(message.Chat.ID)
}

func HandleMainMenu(bot *tgbotapi.BotAPI, message *tgbotapi.Message, stateManager *states.StateManager) {
	HandleStart(bot, message, stateManager)
}
