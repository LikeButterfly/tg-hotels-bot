package keyboards

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	rapidapi_models "tg-hotels-bot/internal/models/rapidapi"
)

// Создаёт клавиатуру для выбора города
func CreateCitiesMarkup(cities []rapidapi_models.City) *tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton

	for _, city := range cities {
		button := tgbotapi.NewInlineKeyboardButtonData(
			city.Name,
			"search_in_city"+strconv.Itoa(city.ID),
		)
		rows = append(rows, tgbotapi.NewInlineKeyboardRow(button))
	}

	markup := tgbotapi.NewInlineKeyboardMarkup(rows...)
	return &markup
}
