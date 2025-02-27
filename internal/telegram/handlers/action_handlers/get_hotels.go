package action_handlers

import (
	"log"
	"time"

	"tg-hotels-bot/internal/config"
	"tg-hotels-bot/internal/database"
	"tg-hotels-bot/internal/photos"
	"tg-hotels-bot/internal/states"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Обрабатывает команды "/lowprice", "/highprice", "/bestdeal"
func DefineState(cfg *config.Config, bot *tgbotapi.BotAPI, message *tgbotapi.Message, stateManager *states.StateManager, userData map[int64]map[string]string) {
	command := message.Text[1:] // Убираем "/"
	handleSearchCommand(cfg, bot, message, command, stateManager, userData)
}

// Обрабатывает кнопку "Недорогие отели"
func ShowLowprice(cfg *config.Config, bot *tgbotapi.BotAPI, message *tgbotapi.Message, stateManager *states.StateManager, userData map[int64]map[string]string) {
	command := "lowprice"
	handleSearchCommand(cfg, bot, message, command, stateManager, userData)
}

// Выполняет общие действия для поиска отелей
func handleSearchCommand(cfg *config.Config, bot *tgbotapi.BotAPI, message *tgbotapi.Message, command string, stateManager *states.StateManager, userData map[int64]map[string]string) {
	chatID := message.Chat.ID

	// Обновляем состояние пользователя
	stateManager.SetState(chatID, states.WaitCityName)

	// Записываем команду в историю
	registerCommandInDB(cfg, chatID, command, userData)

	// Отправляем фото + запрос на ввод города
	sendCityRequestWithPhoto(bot, chatID, command)
}

// Отправляет фото с запросом города
func sendCityRequestWithPhoto(bot *tgbotapi.BotAPI, chatID int64, command string) {
	photoURL := photos.GetPhotoByCommand(command)

	msg := tgbotapi.NewPhoto(chatID, tgbotapi.FileURL(photoURL))
	msg.Caption = "<b>Отправьте боту город для поиска</b>"
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	bot.Send(msg)
}

// Записывает команду в историю
func registerCommandInDB(cfg *config.Config, chatID int64, command string, userData map[int64]map[string]string) {
	callTime := time.Now()

	if _, exists := userData[chatID]; !exists {
		userData[chatID] = make(map[string]string)
	}
	userData[chatID]["command_type"] = command
	userData[chatID]["command_call_time"] = callTime.Format(time.RFC3339)

	// Добавляем в базу
	err := database.AddCommandToHistory(cfg, command, callTime, chatID)
	if err != nil {
		log.Printf("Ошибка записи команды в историю: %v", err)
	}
}
