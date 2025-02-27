package hotels_handlers

import (
	"log"
	"time"

	"tg-hotels-bot/internal/config"
	"tg-hotels-bot/internal/database"
	"tg-hotels-bot/internal/rapidapi/create_messages"
	"tg-hotels-bot/internal/states"
	"tg-hotels-bot/internal/telegram/handlers/default_handlers"
	"tg-hotels-bot/internal/utils/work_with_messages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Обрабатывает ввод города от пользователя
func GetCitiesByName(cfg *config.Config, bot *tgbotapi.BotAPI, message *tgbotapi.Message, stateManager *states.StateManager) {
	city := message.Text

	// Добавляем город в историю MongoDB
	AddCityNameToDB(cfg, bot, message.Chat.ID, city)

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

// Добавляет город в историю пользователя
func AddCityNameToDB(cfg *config.Config, bot *tgbotapi.BotAPI, chatID int64, cityName string) {
	// Получаем время команды из истории состояний
	callTime := time.Now() // Если в будущем нужно будет сохранять время в StateManager

	err := database.AddCityToHistory(cfg, cityName, callTime, chatID)
	if err != nil {
		log.Printf("Ошибка добавления города в историю: %v", err)
	}
}
