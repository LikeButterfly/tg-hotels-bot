package create_messages

import (
	"fmt"
	inline_keyboards "tg-hotels-bot/src/keyboards/inline"
	"tg-hotels-bot/src/rapidapi/parse_responses"
	"tg-hotels-bot/src/utils"
)

// CreateCitiesMessage создаёт сообщение с найденными городами
func CreateCitiesMessage(city string) (utils.CitiesMessage, error) {
	found, err := parse_responses.FindCities(city)
	if err != nil {
		return utils.CitiesMessage{}, err
	}

	var text string
	if len(found) == 1 {
		for cityName := range found {
			text = fmt.Sprintf("<b>Искать в городе %s?</b>", cityName)
		}
	} else {
		text = "<b>Пожалуйста, уточните город</b>"
	}

	buttons := inline_keyboards.CreateCitiesMarkup(found)
	return utils.CitiesMessage{Message: text, Buttons: buttons}, nil
}
