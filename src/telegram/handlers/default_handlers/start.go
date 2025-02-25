package default_handlers

import (
	// inline_keyboards "tg-hotels-bot/src/keyboards/inline"
	reply_keyboards "tg-hotels-bot/src/keyboards/reply"
	"tg-hotels-bot/src/photos"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleStart(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	photo := tgbotapi.NewPhotoShare(message.Chat.ID, photos.Photos["home_menu"])
	photo.Caption = "<b>Выберите действие</b>"
	photo.ParseMode = "HTML"
	photo.ReplyMarkup = reply_keyboards.HomeMenuKeyboard()
	bot.Send(photo)
}

func HandleMainMenu(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	HandleStart(bot, message)
}
