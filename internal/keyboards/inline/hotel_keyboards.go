package keyboards

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tg-hotels-bot/internal/models"
)

func CreateHotelKeyboard(info models.HotelInfo) *tgbotapi.InlineKeyboardMarkup {
	hotelID := info.ID

	bookingButton := tgbotapi.NewInlineKeyboardButtonURL(
		"Забронировать",
		fmt.Sprintf("https://hotels.com/ho%s", hotelID),
	)

	mapsButton := tgbotapi.NewInlineKeyboardButtonData(
		"На карте",
		fmt.Sprintf("get_hotel_map%f/%f", info.Latitude, info.Longitude), // TODO check
	)

	photosButton := tgbotapi.NewInlineKeyboardButtonData(
		"Фото",
		fmt.Sprintf("get_hotel_photos%s", hotelID),
	)

	// favoriteButton := tgbotapi.NewInlineKeyboardButtonData(
	// 	"Добавить в избранное",
	// 	"add_to_favorites",
	// )

	markup := tgbotapi.NewInlineKeyboardMarkup(
		[]tgbotapi.InlineKeyboardButton{bookingButton},
		[]tgbotapi.InlineKeyboardButton{mapsButton},
		[]tgbotapi.InlineKeyboardButton{photosButton},
		// []tgbotapi.InlineKeyboardButton{favoriteButton},
	)

	return &markup
}
