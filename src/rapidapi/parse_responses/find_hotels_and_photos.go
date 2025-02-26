package parse_responses

import (
	"errors"
	"math"
	"strconv"
	"time"
)

// HotelInfo содержит информацию об отеле
type HotelInfo struct {
	ID                 string
	Name               string
	Stars              int
	Address            string
	DistanceFromCenter float64
	TotalCost          float64
	CostByNight        float64
	Photo              string
	Latitude           float64
	Longitude          float64
}

// GetHotelsInfo ищет отели по параметрам
func GetHotelsInfo(data map[string]any, page int) ([]HotelInfo, error) {
	command, _ := data["command_type"].(string)
	dateIn, _ := data["date_in"].(time.Time)
	dateOut, _ := data["date_out"].(time.Time)

	hotelsDict, err := getHotelsDict(command, data, page)
	if err != nil {
		return nil, err
	}

	if errMsg, exists := hotelsDict["error"].(string); exists {
		return nil, errors.New(errMsg)
	}

	// Парсим отели
	results, err := getResults(hotelsDict)
	if err != nil {
		return nil, err
	}

	return parseHotelsInfo(results, dateIn, dateOut), nil
}

// parseHotelsInfo парсит список отелей
func parseHotelsInfo(results []map[string]any, dateIn, dateOut time.Time) []HotelInfo {
	var hotels []HotelInfo
	daysIn := int(dateOut.Sub(dateIn).Hours() / 24)

	for _, result := range results {
		name, _ := result["name"].(string)
		stars, _ := result["reviews"].(map[string]any)["score"].(float64)
		hotelID, _ := result["id"].(string)
		highResLink := "https://ya.ru/"

		// Достаем расстояние до центра
		distanceStr, _ := result["destinationInfo"].(map[string]any)["distanceFromDestination"].(map[string]any)["value"].(string)
		distance := distanceStrToFloatInKm(distanceStr)

		// Цена за ночь и общая стоимость
		priceByNight, _ := result["price"].(map[string]any)["lead"].(map[string]any)["amount"].(float64)
		totalPrice := priceByNight * float64(daysIn)

		// Координаты
		lat, _ := result["mapMarker"].(map[string]any)["latLong"].(map[string]any)["latitude"].(float64)
		lng, _ := result["mapMarker"].(map[string]any)["latLong"].(map[string]any)["longitude"].(float64)

		hotels = append(hotels, HotelInfo{
			ID:                 hotelID,
			Name:               name,
			Stars:              int(stars),
			Address:            "Dummy Address",
			DistanceFromCenter: distance,
			TotalCost:          totalPrice,
			CostByNight:        priceByNight,
			Photo:              highResLink,
			Latitude:           lat,
			Longitude:          lng,
		})
	}

	return hotels
}

func distanceStrToFloatInKm(strDistance string) float64 {
	distance, _ := strconv.ParseFloat(strDistance, 64)
	return math.Round(distance*1.609*100) / 100
}
