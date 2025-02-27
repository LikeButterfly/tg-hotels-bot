package hotels_handlers

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tg-hotels-bot/internal/states"
)

// Обрабатывает выбор города пользователем
func SetCityID(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery, stateManager *states.StateManager, userData map[int64]map[string]string) {
	chatID := callback.Message.Chat.ID

	// Извлекаем city_id из callback-данных
	cityID := strings.TrimPrefix(callback.Data, "search_in_city")
	if cityID == "" {
		log.Println("Ошибка: city_id пуст")
		return
	}

	// Обновляем состояние пользователя
	if _, exists := userData[chatID]; !exists {
		userData[chatID] = make(map[string]string)
	}
	userData[chatID]["city_id"] = cityID

	// Отправляем подтверждение выбора города
	callbackConfig := tgbotapi.NewCallback(callback.ID, "Город выбран")
	bot.Send(callbackConfig)

	editMsg := tgbotapi.NewEditMessageText(chatID, callback.Message.MessageID, "<b>Город выбран!</b>")
	editMsg.ParseMode = "HTML"
	bot.Send(editMsg)

	// Проверяем команду пользователя
	commandType, exists := userData[chatID]["command_type"]
	if !exists {
		log.Println("Ошибка: command_type отсутствует в userData")
	}

	// Если команда bestdeal (поиск по цене и расстоянию), переход к выбору диапазона цен
	if commandType == "bestdeal" {
		// TODO: start_select_price_range(bot, chatID, stateManager, userData)
		return
	}

	// Если обычный поиск, переходим к выбору даты заезда
	StartSelectDateIn(bot, chatID, stateManager)
}
