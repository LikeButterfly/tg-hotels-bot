package hotels_handlers

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tg-hotels-bot/internal/models"
	"tg-hotels-bot/internal/states"
	misc_utils "tg-hotels-bot/internal/utils/misc"
)

// Начинает процесс выбора даты заезда
func StartSelectDateIn(bot *tgbotapi.BotAPI, chatID int64, stateManager *states.StateManager) {
	msg := tgbotapi.NewMessage(chatID, "Пожалуйста, введите дату заезда в формате ДД.ММ.ГГГГ:")
	msg.ParseMode = "HTML"
	bot.Send(msg)

	stateManager.SetState(chatID, states.SelectDateIn)
}

// Обрабатывает введённую пользователем дату заезда
func SelectDateIn(bot *tgbotapi.BotAPI, message *tgbotapi.Message, stateManager *states.StateManager, usersData map[int64]*models.UserData) {
	chatID := message.Chat.ID
	dateStr := message.Text

	// Парсим дату в формате ДД.ММ.ГГГГ
	dateIn, err := time.Parse("02.01.2006", dateStr)
	if err != nil {
		errMsg := tgbotapi.NewMessage(chatID, "Неверный формат даты. Введите в формате ДД.ММ.ГГГГ:")
		bot.Send(errMsg)
		return
	}

	// Проверяем, что дата не в прошлом
	today := time.Now().Truncate(24 * time.Hour)
	if dateIn.Before(today) {
		errMsg := tgbotapi.NewMessage(chatID, "Дата заезда не может быть в прошлом. Попробуйте снова (ДД.ММ.ГГГГ):")
		bot.Send(errMsg)
		return
	}

	// Сохраняем дату заезда в usersData (под ключом "date_in")
	usersData[chatID].DateIn = dateIn

	// Переходим к запросу даты выезда
	StartSelectDateOut(bot, chatID, dateIn, stateManager)
}

// Начинает процесс выбора даты выезда
func StartSelectDateOut(bot *tgbotapi.BotAPI, chatID int64, dateIn time.Time, stateManager *states.StateManager) {
	msg := tgbotapi.NewMessage(chatID, "Введите дату выезда в формате ДД.ММ.ГГГГ:")
	msg.ParseMode = "HTML"
	bot.Send(msg)

	stateManager.SetState(chatID, states.SelectDateOut)
}

// Обрабатывает введённую пользователем дату выезда
func SelectDateOut(bot *tgbotapi.BotAPI, message *tgbotapi.Message, stateManager *states.StateManager, usersData map[int64]*models.UserData) {
	chatID := message.Chat.ID
	dateStr := message.Text

	dateOut, err := time.Parse("02.01.2006", dateStr)
	if err != nil {
		errMsg := tgbotapi.NewMessage(chatID, "Неверный формат даты. Введите в формате ДД.ММ.ГГГГ:")
		bot.Send(errMsg)
		return
	}

	dateIn := usersData[chatID].DateIn

	// Проверяем, что дата выезда не раньше заезда
	if dateOut.Before(dateIn) {
		errMsg := tgbotapi.NewMessage(chatID, "Дата выезда не может быть раньше даты заезда. Попробуйте снова (ДД.ММ.ГГГГ):")
		bot.Send(errMsg)
		return
	}

	// Сохраняем дату выезда в usersData (под ключом "date_out")
	usersData[chatID].DateOut = dateOut

	// Спрашиваем подтверждение
	finalMsg := fmt.Sprintf("<b>Выбрано:</b>\nЗаезд: %s\nВыезд: %s\nВсе верно?", dateIn, dateOut)
	msg := tgbotapi.NewMessage(chatID, finalMsg)
	msg.ParseMode = "HTML"
	msg.ReplyMarkup = misc_utils.IsCorrectMarkup("city_info")
	if _, err := bot.Send(msg); err != nil {
		log.Println("Ошибка отправки сообщения: ", err)
	}

}
