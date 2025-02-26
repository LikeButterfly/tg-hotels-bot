package action_handlers

import (
	"log"
	"time"

	"tg-hotels-bot/internal/database"
	"tg-hotels-bot/internal/photos"
	"tg-hotels-bot/internal/states"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// DefineState обрабатывает команды "/lowprice", "/highprice", "/bestdeal"
func DefineState(bot *tgbotapi.BotAPI, message *tgbotapi.Message, stateManager *states.StateManager, userData map[int64]map[string]string) {
	command := message.Text[1:] // Убираем "/"
	handleSearchCommand(bot, message, command, stateManager, userData)
}

// ShowLowprice обрабатывает кнопку "Недорогие отели"
func ShowLowprice(bot *tgbotapi.BotAPI, message *tgbotapi.Message, stateManager *states.StateManager, userData map[int64]map[string]string) {
	command := "lowprice"
	handleSearchCommand(bot, message, command, stateManager, userData)
}

// handleSearchCommand выполняет общие действия для поиска отелей
func handleSearchCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message, command string, stateManager *states.StateManager, userData map[int64]map[string]string) {
	chatID := message.Chat.ID

	// Обновляем состояние пользователя
	stateManager.SetState(chatID, states.WaitCityName)

	// Записываем команду в историю
	registerCommandInDB(chatID, command, userData)

	// Отправляем фото + запрос на ввод города
	sendCityRequestWithPhoto(bot, chatID, command)
}

// sendCityRequestWithPhoto отправляет фото с запросом города
func sendCityRequestWithPhoto(bot *tgbotapi.BotAPI, chatID int64, command string) {
	photoURL := photos.GetPhotoByCommand(command)

	msg := tgbotapi.NewPhotoShare(chatID, photoURL)
	msg.Caption = "<b>Отправьте боту город для поиска</b>"
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	bot.Send(msg)
}

// registerCommandInDB записывает команду в историю
func registerCommandInDB(chatID int64, command string, userData map[int64]map[string]string) {
	callTime := time.Now()

	if _, exists := userData[chatID]; !exists {
		userData[chatID] = make(map[string]string)
	}
	userData[chatID]["command_type"] = command
	userData[chatID]["command_call_time"] = callTime.Format(time.RFC3339)

	// Добавляем в базу
	err := database.AddCommandToHistory(command, callTime, chatID)
	if err != nil {
		log.Printf("Ошибка записи команды в историю: %v", err)
	}
}
