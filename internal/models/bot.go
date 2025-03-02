package models

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tg-hotels-bot/internal/states"
)

type UserData struct {
	CityID          string
	DateIn          string
	DateOut         string
	CommandType     string
	CommandCallTime string
}

type Bot struct {
	BotAPI       *tgbotapi.BotAPI
	StateManager *states.StateManager
	UsersData    map[int64]*UserData
}
