package keyboards

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CustomCalendar struct {
	CurrentStep string
	MinDate     time.Time
}

func NewCustomCalendar(minDate time.Time) *CustomCalendar {
	return &CustomCalendar{
		CurrentStep: "y", // Начинаем с выбора года
		MinDate:     minDate,
	}
}

var CUSTOM_STEPS = map[string]string{
	"y": "год",
	"m": "месяц",
	"d": "день",
}

func CreateCalendar(minDate time.Time) (tgbotapi.InlineKeyboardMarkup, string) {
	calendar := NewCustomCalendar(minDate)

	prevButton := tgbotapi.NewInlineKeyboardButtonData("⬅️", "prev_"+calendar.CurrentStep)
	nextButton := tgbotapi.NewInlineKeyboardButtonData("➡️", "next_"+calendar.CurrentStep)
	selectButton := tgbotapi.NewInlineKeyboardButtonData(
		fmt.Sprintf("Выбрать %s", CUSTOM_STEPS[calendar.CurrentStep]), "select_"+calendar.CurrentStep,
	)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		[]tgbotapi.InlineKeyboardButton{prevButton, nextButton},
		[]tgbotapi.InlineKeyboardButton{selectButton},
	)

	return keyboard, CUSTOM_STEPS[calendar.CurrentStep]
}

// Обрабатывает callback-данные от inline-календаря
func ProcessCalendar(callbackData string, minDate time.Time) (time.Time, *tgbotapi.InlineKeyboardMarkup, string) {
	var selectedDate time.Time
	var step string

	if callbackData == "select_y" {
		selectedDate = time.Date(minDate.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
		step = "m"
	} else if callbackData == "select_m" {
		selectedDate = time.Date(minDate.Year(), minDate.Month(), 1, 0, 0, 0, 0, time.UTC)
		step = "d"
	} else if callbackData == "select_d" {
		selectedDate = minDate
		return selectedDate, nil, ""
	} else {
		return time.Time{}, nil, ""
	}

	keyboard, _ := CreateCalendar(selectedDate)
	return time.Time{}, &keyboard, step
}
