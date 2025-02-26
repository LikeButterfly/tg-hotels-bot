package parse_responses

import (
	"errors"
	"tg-hotels-bot/internal/rapidapi/rapidapi_requests"
	"time"
)

// getHotelsDict получает список отелей от RapidAPI
func getHotelsDict(command string, data map[string]any, page int) (map[string]any, error) {
	sortBy := "PRICE_LOW_TO_HIGH"
	if command == "highprice" {
		sortBy = "PRICE_RELEVANT"
	}

	destinationID, _ := data["city_id"].(string)
	dateIn, _ := data["date_in"].(time.Time)
	dateOut, _ := data["date_out"].(time.Time)

	hotelsDict, err := rapidapi_requests.GetHotelsJSON(destinationID, dateIn, dateOut, sortBy, page)
	if err != nil {
		return nil, err
	}

	return hotelsDict, nil
}

// getResults парсит JSON-ответ API и возвращает список отелей
func getResults(hotels map[string]any) ([]map[string]any, error) {
	properties, ok := hotels["data"].(map[string]any)["propertySearch"].(map[string]any)["properties"].([]any)
	if !ok {
		return nil, errors.New("ошибка парсинга данных")
	}

	var results []map[string]any
	for _, prop := range properties {
		if hotel, ok := prop.(map[string]any); ok {
			results = append(results, hotel)
		}
	}

	return results, nil
}
