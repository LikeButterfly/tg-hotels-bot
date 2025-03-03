package models

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tg-hotels-bot/internal/states"
)

type UserData struct {
	CityID          string
	DateIn          time.Time
	DateOut         time.Time
	CommandType     CommandType
	CommandCallTime string
	HotelsInfo      []HotelInfo
	HotelIndex      int
}

type Bot struct {
	BotAPI       *tgbotapi.BotAPI
	StateManager *states.StateManager
	UsersData    map[int64]*UserData
}
