package utils

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type KM float64
type USD float64
type Link string
type PhotoID string
type ID int64
type Degrees float64
type Latitude Degrees
type Longitude Degrees

type CitiesMessage struct {
	Message string
	Buttons *tgbotapi.InlineKeyboardMarkup
}

type CalendarMarkupAndStep struct {
	Calendar *tgbotapi.InlineKeyboardMarkup
	DateType string
}

type HotelInfo struct {
	HotelID            ID
	Name               string
	Stars              int
	Address            string
	DistanceFromCenter KM
	TotalCost          USD
	CostByNight        USD
	Photo              Link
	Coordinates        struct {
		Lat Latitude
		Lng Longitude
	}
}

type HotelMessage struct {
	Text    string
	Photo   Link
	Buttons *tgbotapi.InlineKeyboardMarkup
}
