package rapidapi_requests

import (
	"encoding/json"
	"errors"
	"os"

	rapidapi_models "tg-hotels-bot/internal/models/rapidapi"
)

// Запрос для поиска городов
func GetCitiesJSON(city string) (*rapidapi_models.CitiesResponse, error) {
	url := "https://hotels4.p.rapidapi.com/locations/v3/search"

	// FIXME?
	locale := os.Getenv("LOCALE_IDENTIFIER")
	langID := os.Getenv("LANGUAGE_IDENTIFIER")
	siteID := os.Getenv("SITE_ID")

	queryParams := map[string]string{
		"q":      city,
		"locale": locale,
		"langid": langID,
		"siteid": siteID,
	}

	responseBytes, err := RequestToAPI(url, queryParams)
	if err != nil {
		return nil, errors.New("ошибка запроса к API: " + err.Error())
	}

	if responseBytes == nil {
		return nil, errors.New("пустой ответ от API")
	}

	data, err := json.Marshal(responseBytes)
	if err != nil {
		return nil, errors.New("ошибка маршаллинга ответа: " + err.Error())
	}

	var citiesResp rapidapi_models.CitiesResponse
	if err := json.Unmarshal(data, &citiesResp); err != nil {
		return nil, errors.New("ошибка парсинга ответа: " + err.Error())
	}

	return &citiesResp, nil
}
