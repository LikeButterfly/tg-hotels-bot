package parse_responses

import (
	"encoding/json"
	"errors"

	"tg-hotels-bot/internal/models"
	"tg-hotels-bot/internal/rapidapi/rapidapi_requests"
)

type HotelsAPIResponse struct {
	Data struct {
		PropertySearch struct {
			Properties []HotelProperty `json:"properties"`
		} `json:"propertySearch"`
	} `json:"data"`
}

type HotelProperty struct {
	ID              string                       `json:"id"`
	Name            string                       `json:"name"`
	Availability    HotelPropertyAvailability    `json:"availability"`
	PropertyImage   HotelPropertyImage           `json:"propertyImage"`
	DestinationInfo HotelPropertyDestinationInfo `json:"destinationInfo"`
	Reviews         HotelPropertyReviews         `json:"reviews"`
	Price           HotelPropertyPrice           `json:"price"`
	MapMarker       HotelPropertyMapMarker       `json:"mapMarker"`
}

type HotelPropertyAvailability struct {
	Available    bool `json:"available"`
	MinRoomsLeft int  `json:"minRoomsLeft"`
}

type HotelPropertyImage struct {
	HotelImage struct {
		ImageURL string `json:"url"`
	} `json:"image"`
}

type HotelPropertyDestinationInfo struct {
	DistanceFromDestination HotelDistance `json:"distanceFromDestination"`
	RegionID                string        `json:"regionId"`
}

type HotelDistance struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type HotelPropertyReviews struct {
	Score float64 `json:"score"`
	Total int     `json:"total"`
}

type HotelPropertyPrice struct {
	Lead PriceLead `json:"lead"`
}

type PriceLead struct {
	Amount float64 `json:"amount"`
}

type HotelPropertyMapMarker struct {
	LatLong Coordinates `json:"latLong"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Получает список отелей от RapidAPI
func getHotelsDict(data *models.UserData) (map[string]any, error) {
	sortBy := "PRICE_LOW_TO_HIGH"
	if data.CommandType == "highprice" {
		sortBy = "PRICE_RELEVANT"
	}

	destinationID := data.CityID

	hotelsDict, err := rapidapi_requests.GetHotelsJSON(destinationID, data.DateIn, data.DateOut, sortBy)
	if err != nil {
		return nil, err
	}

	return hotelsDict, nil
}

func getResults(hotels map[string]any) ([]HotelProperty, error) {
	// Преобразуем map в JSON
	jsonData, err := json.Marshal(hotels)
	if err != nil {
		return nil, errors.New("ошибка сериализации данных: " + err.Error())
	}

	var apiResp HotelsAPIResponse
	if err := json.Unmarshal(jsonData, &apiResp); err != nil {
		return nil, errors.New("ошибка десериализации данных: " + err.Error())
	}

	return apiResp.Data.PropertySearch.Properties, nil
}
