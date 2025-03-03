package parse_responses

import (
	"math"
	"time"

	"tg-hotels-bot/internal/models"
)

// Ищет отели по параметрам
func GetHotelsInfo(data *models.UserData) ([]models.HotelInfo, error) {
	hotelsDict, err := getHotelsDict(data)
	if err != nil {
		return nil, err
	}

	// Парсим отели
	results, err := getResults(hotelsDict)
	if err != nil {
		return nil, err
	}

	return parseHotelsInfo(results, data.DateIn, data.DateOut), nil
}

// Парсит список отелей
func parseHotelsInfo(results []HotelProperty, dateIn, dateOut time.Time) []models.HotelInfo {
	var hotels []models.HotelInfo
	daysIn := int(dateOut.Sub(dateIn).Hours() / 24)

	for _, result := range results {
		hotelID := result.ID
		name := result.Name
		stars := int(result.Reviews.Score)

		// FIXME
		address := "Dummy Address"
		// highResLink := "https://ya.ru/"
		highResLink := result.PropertyImage.HotelImage.ImageURL

		// Дистанция до центра: если unit = "MILE", переводим в км
		distance := result.DestinationInfo.DistanceFromDestination.Value
		if result.DestinationInfo.DistanceFromDestination.Unit == "MILE" {
			distance *= 1.609
		}

		// Цена за ночь и общая стоимость
		priceByNight := result.Price.Lead.Amount     // FIXME
		totalPrice := priceByNight * float64(daysIn) // FIXME

		// Координаты
		lat := result.MapMarker.LatLong.Latitude
		lng := result.MapMarker.LatLong.Longitude

		hotels = append(hotels, models.HotelInfo{
			ID:                 hotelID,
			Name:               name,
			Stars:              stars,
			Address:            address,
			DistanceFromCenter: math.Round(distance*100) / 100,
			TotalCost:          totalPrice,
			CostByNight:        priceByNight,
			Photo:              highResLink,
			Latitude:           lat,
			Longitude:          lng,
		})
	}

	// fmt.Println("--------------++++++++++++")
	// fmt.Println(hotels[:1])
	// fmt.Println("--------------++++++++++++")

	return hotels
}
