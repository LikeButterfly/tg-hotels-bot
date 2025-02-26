package hotels_handlers

import (
	"log"
	"tg-hotels-bot/src/database"
	"tg-hotels-bot/src/rapidapi/create_messages"
	"tg-hotels-bot/src/states"
	"tg-hotels-bot/src/telegram/handlers/default_handlers"
	"tg-hotels-bot/src/utils/work_with_messages"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// GetCitiesByName обрабатывает ввод города от пользователя
func GetCitiesByName(bot *tgbotapi.BotAPI, message *tgbotapi.Message, stateManager *states.StateManager) {
	city := message.Text

	// Добавляем город в историю MongoDB
	AddCityNameToDB(bot, message.Chat.ID, city)

	// Отправляем сообщение "Ожидание" и получаем ID сообщений для удаления
	textMsgID, stickerMsgID := work_with_messages.SendWaitingMessage(bot, message.Chat.ID)

	// Получаем список городов
	citiesMessage, err := create_messages.CreateCitiesMessage(city)
	work_with_messages.DeleteWaitingMessages(bot, message.Chat.ID, textMsgID, stickerMsgID)

	if err != nil {
		default_handlers.FinishWithError(bot, message.Chat.ID, "cities_not_found", stateManager)
		return
	}

	// Отправляем пользователю список городов с кнопками
	msg := tgbotapi.NewMessage(message.Chat.ID, citiesMessage.Message)
	msg.ReplyMarkup = citiesMessage.Buttons
	bot.Send(msg)

	// Переключаем состояние на выбор города
	stateManager.SetState(message.Chat.ID, states.SelectCity)
}

// AddCityNameToDB добавляет город в историю пользователя
func AddCityNameToDB(bot *tgbotapi.BotAPI, chatID int64, cityName string) {
	// Получаем время команды из истории состояний
	callTime := time.Now() // Если в будущем нужно будет сохранять время в StateManager

	err := database.AddCityToHistory(cityName, callTime, chatID)
	if err != nil {
		log.Printf("Ошибка добавления города в историю: %v", err)
	}
}
