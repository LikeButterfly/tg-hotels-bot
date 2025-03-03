package create_messages

import (
	"fmt"
	"strings"

	inline_keyboards "tg-hotels-bot/internal/keyboards/inline"
	"tg-hotels-bot/internal/models"
	"tg-hotels-bot/internal/utils"
)

// Создаёт сообщение с информацией об отеле
func CreateHotelMessage(hotelInfo models.HotelInfo) utils.HotelMessage {
	text := fmt.Sprintf(
		"<b>%s</b>\n%s"+
			"\tАдрес: %s\n"+
			"\tРасстояние до центра: %.2f км\n"+
			"\tСтоимость: %.2f $\n"+
			"\tСтоимость за ночь: %.2f $\n",
		hotelInfo.Name,
		getStarsString(hotelInfo.Stars),
		hotelInfo.Address,
		hotelInfo.DistanceFromCenter,
		hotelInfo.TotalCost,
		hotelInfo.CostByNight,
	)

	buttons := inline_keyboards.CreateHotelKeyboard(hotelInfo)
	return utils.HotelMessage{Text: text, Photo: hotelInfo.Photo, Buttons: buttons}
}

// Возвращает строку с эмодзи звёзд
func getStarsString(stars int) string {
	if stars < 1 {
		return ""
	}
	return fmt.Sprintf("\t%s\n", strings.Repeat("⭐️", stars))
}
