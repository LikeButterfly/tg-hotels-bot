package rapidapi_requests

import (
	"errors"
	"os"
	"strconv"
	"time"
)

type Child struct {
	Age int `json:"age"`
}

type Room struct {
	Adults   int     `json:"adults"`
	Children []Child `json:"children"`
}

type Date struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type Destination struct {
	RegionID string `json:"regionId"`
}

type HotelsPayload struct {
	Currency             string      `json:"currency"`
	Eapid                string      `json:"eapid"`
	Locale               string      `json:"locale"`
	SiteID               int         `json:"siteId"`
	Destination          Destination `json:"destination"`
	CheckInDate          Date        `json:"checkInDate"`
	CheckOutDate         Date        `json:"checkOutDate"`
	Rooms                []Room      `json:"rooms"`
	ResultsStartingIndex int         `json:"resultsStartingIndex"`
	ResultsSize          int         `json:"resultsSize"`
	Sort                 string      `json:"sort"`
}

// Запрос для поиска отелей
func GetHotelsJSON(destinationID string, dateIn, dateOut time.Time, sortBy string) (map[string]any, error) {
	url := "https://hotels4.p.rapidapi.com/properties/v2/list"

	eapid := os.Getenv("EAPID")
	siteIDStr := os.Getenv("SITE_ID")

	siteIDInt, err := strconv.Atoi(siteIDStr)
	if err != nil {
		return nil, errors.New("Ошибка преобразования SITE_ID в число: " + err.Error())
	}
	siteID := siteIDInt

	payload := HotelsPayload{
		Currency: "USD",
		Eapid:    eapid,
		Locale:   "en_US",
		SiteID:   siteID,
		Destination: Destination{
			RegionID: destinationID,
		},
		CheckInDate: Date{
			Day:   dateIn.Day(),
			Month: int(dateIn.Month()),
			Year:  dateIn.Year(),
		},
		CheckOutDate: Date{
			Day:   dateOut.Day(),
			Month: int(dateOut.Month()),
			Year:  dateOut.Year(),
		},
		Rooms: []Room{
			{
				Adults:   1,
				Children: []Child{},
			},
		},
		ResultsStartingIndex: 0,
		ResultsSize:          15,
		Sort:                 sortBy,
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
