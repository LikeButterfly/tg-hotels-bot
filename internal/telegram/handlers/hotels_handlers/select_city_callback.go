package hotels_handlers

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tg-hotels-bot/internal/models"
	"tg-hotels-bot/internal/states"
)

// Обрабатывает выбор города пользователем
func SetCityID(
	bot *tgbotapi.BotAPI,
	callback *tgbotapi.CallbackQuery,
	stateManager *states.StateManager,
	usersData map[int64]*models.UserData,
) {
	chatID := callback.Message.Chat.ID

	// Извлекаем city_id из callback-данных
	cityID := strings.TrimPrefix(callback.Data, "search_in_city")
	if cityID == "" {
		log.Println("Ошибка: city_id пуст")
		// TODO обработать...
		return
	}

	// Обновляем состояние пользователя
	if _, exists := usersData[chatID]; !exists {
		usersData[chatID] = &models.UserData{}
	}
	usersData[chatID].CityID = cityID

	// Отправляем подтверждение выбора города
	editMsg := tgbotapi.NewEditMessageText(chatID, callback.Message.MessageID, "<b>Город выбран!</b>")
	editMsg.ParseMode = "HTML"
	if _, err := bot.Send(editMsg); err != nil {
		// TODO обработать ошибку
		return
	}

	// Проверяем команду пользователя
	commandType := usersData[chatID].CommandType
	if commandType == "" {
		log.Println("Ошибка: command_type отсутствует в usersData")
		return
	}

	// Если обычный поиск, переходим к выбору даты заезда
	StartSelectDateIn(bot, chatID, stateManager)
}
