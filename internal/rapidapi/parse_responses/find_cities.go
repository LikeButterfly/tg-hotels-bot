package parse_responses

import (
	"errors"
	"strconv"
	"tg-hotels-bot/internal/rapidapi/rapidapi_requests"
)

// Ищет города по названию
func FindCities(city string) (map[string]int, error) {
	citiesDict, err := rapidapi_requests.GetCitiesJSON(city)
	if err != nil {
		return nil, err
	}

	if errMessage, exists := citiesDict["error"].(string); exists {
		return nil, errors.New(errMessage)
	}

	// Достаём список городов
	citySr, exists := citiesDict["sr"].([]any)
	if !exists || len(citySr) == 0 {
		return nil, errors.New("cities_not_found")
	}

	// Создаём мапу городов
	citiesWithID := make(map[string]int)
	for _, cityEntry := range citySr {
		cityData, ok := cityEntry.(map[string]any)
		if !ok {
			continue
		}
		name, _ := cityData["regionNames"].(map[string]any)["shortName"].(string)
		cityID, _ := cityData["gaiaId"].(string)
		var intCityID int // FIXME ?
		intCityID, err = strconv.Atoi(cityID)

		if err != nil {
			continue
		}
		citiesWithID[name] = intCityID
	}

	return citiesWithID, nil
}
