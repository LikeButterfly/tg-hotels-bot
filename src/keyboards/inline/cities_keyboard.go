package keyboards

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func CreateCitiesMarkup(cities map[string]int) tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton
	cityIndex := 1

	for cityName, cityID := range cities {
		button := tgbotapi.NewInlineKeyboardButtonData(
			cityName,
			"search_in_city"+strconv.Itoa(cityID),
		)
		rows = append(rows, tgbotapi.NewInlineKeyboardRow(button))
		cityIndex++
	}

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}
