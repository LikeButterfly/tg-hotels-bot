package rapidapi_requests

import (
	"errors"
	"os"
)

// Запрос для поиска городов
func GetCitiesJSON(city string) (map[string]any, error) {
	url := "https://hotels4.p.rapidapi.com/locations/v3/search"

	locale := os.Getenv("LOCALE_IDENTIFIER")
	langID := os.Getenv("LANGUAGE_IDENTIFIER")
	siteID := os.Getenv("SITE_ID")

	queryParams := map[string]string{
		"q":      city,
		"locale": locale,
		"langid": langID,
		"siteid": siteID,
	}

	response, err := RequestToAPI(url, queryParams)
	if err != nil {
		return nil, errors.New("ошибка запроса к API: " + err.Error())
	}

	if response == nil {
		return nil, errors.New("пустой ответ от API")
	}

	return response, nil
}
