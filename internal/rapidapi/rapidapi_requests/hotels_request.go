package rapidapi_requests

import (
	"errors"
	"os"
	"time"
)

// Запрос для поиска отелей
func GetHotelsJSON(destinationID string, dateIn, dateOut time.Time, sortBy string, page int) (map[string]any, error) {
	url := "https://hotels4.p.rapidapi.com/properties/v2/list"

	eapid := os.Getenv("EAPID")
	siteID := os.Getenv("SITE_ID")

	payload := map[string]any{
		"currency": "USD",
		"eapid":    eapid,
		"locale":   "en_US",
		"siteId":   siteID,
		"destination": map[string]string{
			"regionId": destinationID,
		},
		"checkInDate": map[string]int{
			"day":   dateIn.Day(),
			"month": int(dateIn.Month()),
			"year":  dateIn.Year(),
		},
		"checkOutDate": map[string]int{
			"day":   dateOut.Day(),
			"month": int(dateOut.Month()),
			"year":  dateOut.Year(),
		},
		"rooms": []map[string]any{
			{"adults": 1, "children": []any{}},
		},
		"resultsStartingIndex": 0,
		"resultsSize":          15,
		"sort":                 sortBy,
	}

	response, err := RequestToAPIWithPayload(url, payload)
	if err != nil {
		return nil, errors.New("ошибка запроса к API: " + err.Error())
	}

	if response == nil {
		return nil, errors.New("пустой ответ от API")
	}

	return response, nil
}
