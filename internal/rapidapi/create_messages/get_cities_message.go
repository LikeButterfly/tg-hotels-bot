package create_messages

import (
	"fmt"

	inline_keyboards "tg-hotels-bot/internal/keyboards/inline"
	"tg-hotels-bot/internal/rapidapi/parse_responses"
	"tg-hotels-bot/internal/utils"
)

// Создаёт сообщение с найденными городами
func CreateCitiesMessage(city string) (utils.CitiesMessage, error) {
	cities, err := parse_responses.FindCities(city)
	if err != nil {
		return utils.CitiesMessage{}, err
	}

	var text string
	if len(cities) == 1 {
		text = fmt.Sprintf("<b>Искать в городе %s?</b>", cities[0].Name)
	} else {
		text = "<b>Пожалуйста, уточните город</b>"
	}

	buttons := inline_keyboards.CreateCitiesMarkup(cities)
	return utils.CitiesMessage{Message: text, Buttons: buttons}, nil
}
