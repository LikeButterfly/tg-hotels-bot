package hotels_handlers

import (
	"fmt"
	"log"
	"time"

	inline_keyboards "tg-hotels-bot/internal/keyboards/inline"
	"tg-hotels-bot/internal/states"
	misc_utils "tg-hotels-bot/internal/utils/misc"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Нначинает процесс выбора даты заезда
func StartSelectDateIn(bot *tgbotapi.BotAPI, chatID int64, stateManager *states.StateManager) {
	keyboard, step := inline_keyboards.CreateCalendar(time.Now()) // Передаём текущую дату как минимум

	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("<b>Укажите %s заезда</b>", step))
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "HTML"
	bot.Send(msg)

	// Устанавливаем состояние выбора даты заезда
	stateManager.SetState(chatID, states.SelectDateIn)
}

// Обрабатывает выбор даты заезда
func SelectDateIn(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery, stateManager *states.StateManager, userData map[int64]time.Time) {
	chatID := callback.Message.Chat.ID

	// Получаем дату из callback-данных
	date, keyboard, step := inline_keyboards.ProcessCalendar(callback.Data, time.Now())

	if date.IsZero() && keyboard != nil {
		editMsg := tgbotapi.NewEditMessageText(chatID, callback.Message.MessageID, fmt.Sprintf("<b>Укажите %s заезда</b>", step))
		editMsg.ReplyMarkup = keyboard
		editMsg.ParseMode = "HTML"
		bot.Send(editMsg)
		return
	}

	userData[chatID] = date
	editMsg := tgbotapi.NewEditMessageText(chatID, callback.Message.MessageID, fmt.Sprintf("<b>Выбрано: %s\nВсе верно?</b>", date.Format("02.01.2006")))
	editMsg.ReplyMarkup = misc_utils.IsCorrectMarkup("date_in")
	editMsg.ParseMode = "HTML"
	bot.Send(editMsg)

	stateManager.SetState(chatID, states.IsDateCorrect)
}

// Начинает процесс выбора даты выезда
func StartSelectDateOut(bot *tgbotapi.BotAPI, chatID int64, dateIn time.Time, stateManager *states.StateManager) {
	keyboard, step := inline_keyboards.CreateCalendar(dateIn) // Передаём дату заезда как минимум

	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("<b>Укажите %s выезда</b>", step))
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "HTML"
	bot.Send(msg)

	stateManager.SetState(chatID, states.SelectDateOut)
}

// Обрабатывает выбор даты выезда
func SelectDateOut(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery, stateManager *states.StateManager, userData map[int64]time.Time) {
	chatID := callback.Message.Chat.ID
	dateIn, exists := userData[chatID]
	if !exists {
		log.Println("Ошибка: отсутствует дата заезда")
		return
	}

	date, keyboard, step := inline_keyboards.ProcessCalendar(callback.Data, dateIn)

	if date.IsZero() && keyboard != nil {
		editMsg := tgbotapi.NewEditMessageText(chatID, callback.Message.MessageID, fmt.Sprintf("<b>Укажите %s выезда</b>", step))
		editMsg.ReplyMarkup = keyboard
		editMsg.ParseMode = "HTML"
		bot.Send(editMsg)
		return
	}

	userData[chatID] = date
	editMsg := tgbotapi.NewEditMessageText(chatID, callback.Message.MessageID, fmt.Sprintf("<b>Выбрано: %s\nВсе верно?</b>", date.Format("02.01.2006")))
	editMsg.ReplyMarkup = misc_utils.IsCorrectMarkup("date_out")
	editMsg.ParseMode = "HTML"
	bot.Send(editMsg)

	stateManager.SetState(chatID, states.IsDateCorrect)
}

// Проверяет правильность выбранных дат
func SendConfirmationDate(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery, stateManager *states.StateManager, userData map[int64]time.Time) {
	chatID := callback.Message.Chat.ID
	data := callback.Data

	dateIn := userData[chatID]
	dateOut := userData[chatID]

	// FIXME - грязно...

	switch data {
	case "date_in_incorrect", "date_out_incorrect":
		if _, err := bot.Send(tgbotapi.NewCallback(callback.ID, "Попробуйте еще раз")); err != nil {
			log.Println("Ошибка ответа на callback:", err)
		}
		if _, err := bot.Send(tgbotapi.NewDeleteMessage(chatID, callback.Message.MessageID)); err != nil {
			log.Println("Ошибка удаления сообщения:", err)
		}
		if data == "date_in_incorrect" {
			StartSelectDateIn(bot, chatID, stateManager)
		} else {
			StartSelectDateOut(bot, chatID, dateIn, stateManager)
		}
		return

	case "date_in_correct":
		if _, err := bot.Send(tgbotapi.NewCallback(callback.ID, "Укажите дату выезда")); err != nil {
			log.Println("Ошибка ответа на callback:", err)
		}
		if _, err := bot.Send(tgbotapi.NewDeleteMessage(chatID, callback.Message.MessageID)); err != nil {
			log.Println("Ошибка удаления сообщения:", err)
		}
		StartSelectDateOut(bot, chatID, dateIn, stateManager)
		return

	case "date_out_correct":
		if _, err := bot.Send(tgbotapi.NewCallback(callback.ID, "Дата выезда указана")); err != nil {
			log.Println("Ошибка ответа на callback:", err)
		}
		if _, err := bot.Send(tgbotapi.NewDeleteMessage(chatID, callback.Message.MessageID)); err != nil {
			log.Println("Ошибка удаления сообщения:", err)
		}

		// Отправляем сообщение с итоговой информацией
		finalMsg := fmt.Sprintf(
			"<b>Дата заезда: </b>%s\n<b>Дата выезда: </b>%s\n\n<b>Все верно?</b>",
			dateIn.Format("02.01.2006"),
			dateOut.Format("02.01.2006"),
		)
		msg := tgbotapi.NewMessage(chatID, finalMsg)
		msg.ReplyMarkup = misc_utils.IsCorrectMarkup("city_info")
		msg.ParseMode = "HTML"
		if _, err := bot.Send(msg); err != nil {
			log.Println("Ошибка отправки сообщения:", err)
		}
		stateManager.SetState(chatID, states.IsInfoCorrect)
	}

}
