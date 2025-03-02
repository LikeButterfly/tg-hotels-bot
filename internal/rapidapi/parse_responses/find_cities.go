package parse_responses

import (
	"errors"
	"fmt"
	"strconv"

	rapidapi_models "tg-hotels-bot/internal/models/rapidapi"
	"tg-hotels-bot/internal/rapidapi/rapidapi_requests"
)

// Ищет города по названию
func FindCities(city string) ([]rapidapi_models.City, error) {
	citiesResp, err := rapidapi_requests.GetCitiesJSON(city)
	if err != nil {
		return nil, err
	}
	fmt.Println(citiesResp) // DELETE ME

	if len(citiesResp.SR) == 0 {
		return nil, errors.New("cities_not_found")
	}

	var cities []rapidapi_models.City
	for _, cityResult := range citiesResp.SR {
		intCityID, err := strconv.Atoi(cityResult.GaiaId)
		if err != nil {
			continue
		}
		cities = append(cities, rapidapi_models.City{
			Name: cityResult.RegionNames.ShortName,
			ID:   intCityID,
		})
	}

	return cities, nil
}
