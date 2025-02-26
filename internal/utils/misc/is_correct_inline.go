package misc

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// IsCorrectMarkup создает inline-клавиатуру "Да/Нет"
func IsCorrectMarkup(textBeforeCorrect string) *tgbotapi.InlineKeyboardMarkup {
	yesButton := tgbotapi.NewInlineKeyboardButtonData("Да", textBeforeCorrect+"_correct")
	noButton := tgbotapi.NewInlineKeyboardButtonData("Нет", textBeforeCorrect+"_incorrect")

	markup := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(yesButton, noButton),
	)

	return &markup
}
