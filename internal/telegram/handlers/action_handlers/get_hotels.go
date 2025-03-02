package action_handlers

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tg-hotels-bot/internal/config"
	"tg-hotels-bot/internal/database"
	"tg-hotels-bot/internal/models"
	"tg-hotels-bot/internal/photos"
	"tg-hotels-bot/internal/states"
)

// Обрабатывает команды "/lowprice", "/highprice", "/bestdeal"
func DefineState(cfg *config.Config, bot *tgbotapi.BotAPI, message *tgbotapi.Message, stateManager *states.StateManager, usersData map[int64]*models.UserData) {
	command := message.Text[1:] // Убираем "/"
	handleSearchCommand(cfg, bot, message, command, stateManager, usersData)
}

// Обрабатывает кнопку "Недорогие отели"
func ShowLowprice(cfg *config.Config, bot *tgbotapi.BotAPI, message *tgbotapi.Message, stateManager *states.StateManager, usersData map[int64]*models.UserData) {
	command := "lowprice"
	handleSearchCommand(cfg, bot, message, command, stateManager, usersData)
}

// Выполняет общие действия для поиска отелей
func handleSearchCommand(cfg *config.Config, bot *tgbotapi.BotAPI, message *tgbotapi.Message, command string, stateManager *states.StateManager, usersData map[int64]*models.UserData) {
	chatID := message.Chat.ID

	// Обновляем состояние пользователя
	stateManager.SetState(chatID, states.WaitCityName)

	// Записываем команду в историю
	registerCommandInDB(cfg, chatID, command, usersData)

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
func registerCommandInDB(cfg *config.Config, chatID int64, command string, usersData map[int64]*models.UserData) {
	callTime := time.Now()

	if _, exists := usersData[chatID]; !exists {
		usersData[chatID] = &models.UserData{}
	}
	usersData[chatID].CommandType = command
	usersData[chatID].CommandCallTime = callTime.Format(time.RFC3339)

	// Добавляем в базу
	err := database.AddCommandToHistory(cfg, command, callTime, chatID)
	if err != nil {
		log.Printf("Ошибка записи команды в историю: %v", err)
	}
}
