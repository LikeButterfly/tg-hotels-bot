package hotels_handlers

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tg-hotels-bot/internal/config"
	"tg-hotels-bot/internal/database"
	reply_keyboards "tg-hotels-bot/internal/keyboards/reply"
	"tg-hotels-bot/internal/models"
	"tg-hotels-bot/internal/rapidapi/create_messages"
	"tg-hotels-bot/internal/rapidapi/parse_responses"
	"tg-hotels-bot/internal/utils"
	"tg-hotels-bot/internal/utils/work_with_messages"
)

func SendFirstHotel(
	cfg *config.Config,
	bot *tgbotapi.BotAPI,
	userData *models.UserData,
	chatId int64,
) {
	// Отправляем сообщение "Ожидание" и получаем ID сообщений для удаления
	textMsgID, stickerMsgID := work_with_messages.SendWaitingMessage(bot, chatId)

	hotels_info, err := parse_responses.GetHotelsInfo(userData)
	if err != nil {
		// TODO выдать юзеру сообщение
		return
	}
	work_with_messages.DeleteWaitingMessages(bot, chatId, textMsgID, stickerMsgID)

	userData.HotelsInfo = hotels_info
	userData.HotelIndex = 1

	first_hotel := hotels_info[0]
	hotel_message := create_messages.CreateHotelMessage(first_hotel)

	keyboard := reply_keyboards.ShowMoreHotelsKeyboard()
	msg := tgbotapi.NewMessage(chatId, "<b>Найденные отели:</b>")
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = keyboard
	if _, err := bot.Send(msg); err != nil {
		// TODO обработать ошибку
		return
	}

	if err := TryingToSendWithPhoto(bot, chatId, hotel_message); err != nil {
		// TODO обработать ошибку
		return
	}

	database.AddHotelToHistory(cfg, hotel_message, time.Now(), chatId)

	//

}

func ShowMoreHotels(
	cfg *config.Config,
	bot *tgbotapi.BotAPI,
	userData *models.UserData,
	chatId int64,
) {
	if userData == nil || userData.HotelsInfo == nil {
		log.Println("Данные отелей не инициализированы")
		return
	}

	if userData.HotelIndex > len(userData.HotelsInfo) {
		log.Println("Отели закончились (Пока без ответа юзеру)")
		return
	}

	hotel := userData.HotelsInfo[userData.HotelIndex]
	hotel_message := create_messages.CreateHotelMessage(hotel)

	userData.HotelIndex += 1

	keyboard := reply_keyboards.ShowMoreHotelsKeyboard()
	msg := tgbotapi.NewMessage(chatId, "<b>Найденные отели:</b>")
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = keyboard
	if _, err := bot.Send(msg); err != nil {
		// TODO обработать ошибку
		return
	}

	if err := TryingToSendWithPhoto(bot, chatId, hotel_message); err != nil {
		// TODO обработать ошибку
		return
	}

	database.AddHotelToHistory(cfg, hotel_message, time.Now(), chatId)
}

func TryingToSendWithPhoto(bot *tgbotapi.BotAPI, chatID int64, hotelMessage utils.HotelMessage) error {
	photoMsg := tgbotapi.NewPhoto(chatID, tgbotapi.FileURL(hotelMessage.Photo))
	photoMsg.Caption = hotelMessage.Text
	photoMsg.ParseMode = "HTML"
	photoMsg.ReplyMarkup = hotelMessage.Buttons

	if _, err := bot.Send(photoMsg); err != nil {
		// Если отправка фото не удалась, отправляем сообщение без фото
		textMsg := tgbotapi.NewMessage(chatID, hotelMessage.Text)
		textMsg.ParseMode = "HTML"
		textMsg.ReplyMarkup = hotelMessage.Buttons
		_, err2 := bot.Send(textMsg)
		return err2
	}
	return nil
}
