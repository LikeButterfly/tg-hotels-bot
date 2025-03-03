package utils

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// type KM float64
// type USD float64
// type Link string
// type PhotoID string
// type ID int64
// type Degrees float64
// type Latitude Degrees
// type Longitude Degrees

// TODO переместить в models

type CitiesMessage struct {
	Message string
	Buttons *tgbotapi.InlineKeyboardMarkup
}

type HotelMessage struct {
	Text    string
	Photo   string
	Buttons *tgbotapi.InlineKeyboardMarkup
}
