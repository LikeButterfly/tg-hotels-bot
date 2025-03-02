package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func HomeMenuKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Справка"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Недорогие отели"),
			// tgbotapi.NewKeyboardButton("Дорогие отели"),
		),
		// tgbotapi.NewKeyboardButtonRow(
		// 	tgbotapi.NewKeyboardButton("Поиск с параметрами"),
		// ),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("История поиска"),
			tgbotapi.NewKeyboardButton("Избранное"),
		),
	)

	keyboard.ResizeKeyboard = true
	return keyboard
}

func ShowMoreHotelsKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Показать еще"),
			tgbotapi.NewKeyboardButton("Главное меню"),
		),
	)
	keyboard.ResizeKeyboard = true
	return keyboard
}
