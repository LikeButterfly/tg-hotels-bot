package keyboards

import (
	"fmt"

	"tg-hotels-bot/internal/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func CreateHotelKeyboard(info utils.HotelInfo) *tgbotapi.InlineKeyboardMarkup {
	hotelID := info.HotelID

	bookingButton := tgbotapi.NewInlineKeyboardButtonURL(
		"Забронировать",
		fmt.Sprintf("https://hotels.com/ho%d", hotelID),
	)

	mapsButton := tgbotapi.NewInlineKeyboardButtonData(
		"На карте",
		fmt.Sprintf("get_hotel_map%f/%f", info.Coordinates.Lat, info.Coordinates.Lng), // TODO check
	)

	photosButton := tgbotapi.NewInlineKeyboardButtonData(
		"Фото",
		fmt.Sprintf("get_hotel_photos%d", hotelID),
	)

	favoriteButton := tgbotapi.NewInlineKeyboardButtonData(
		"Добавить в избранное",
		"add_to_favorites",
	)

	markup := tgbotapi.NewInlineKeyboardMarkup(
		[]tgbotapi.InlineKeyboardButton{bookingButton},
		[]tgbotapi.InlineKeyboardButton{mapsButton},
		[]tgbotapi.InlineKeyboardButton{photosButton},
		[]tgbotapi.InlineKeyboardButton{favoriteButton},
	)

	return &markup
}
