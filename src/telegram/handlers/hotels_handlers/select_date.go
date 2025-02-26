package hotels_handlers

import (
	"fmt"
	"log"
	inline_keyboards "tg-hotels-bot/src/keyboards/inline"
	"tg-hotels-bot/src/states"
	misc_utils "tg-hotels-bot/src/utils/misc"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// StartSelectDateIn начинает процесс выбора даты заезда
func StartSelectDateIn(bot *tgbotapi.BotAPI, chatID int64, stateManager *states.StateManager) {
	keyboard, step := inline_keyboards.CreateCalendar(time.Now()) // Передаём текущую дату как минимум

	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("<b>Укажите %s заезда</b>", step))
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "HTML"
	bot.Send(msg)

	// Устанавливаем состояние выбора даты заезда
	stateManager.SetState(chatID, states.SelectDateIn)
}

// SelectDateIn обрабатывает выбор даты заезда
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

// StartSelectDateOut начинает процесс выбора даты выезда
func StartSelectDateOut(bot *tgbotapi.BotAPI, chatID int64, dateIn time.Time, stateManager *states.StateManager) {
	keyboard, step := inline_keyboards.CreateCalendar(dateIn) // Передаём дату заезда как минимум

	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("<b>Укажите %s выезда</b>", step))
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "HTML"
	bot.Send(msg)

	stateManager.SetState(chatID, states.SelectDateOut)
}

// SelectDateOut обрабатывает выбор даты выезда
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

// SendConfirmationDate проверяет правильность выбранных дат
func SendConfirmationDate(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery, stateManager *states.StateManager, userData map[int64]time.Time) {
	chatID := callback.Message.Chat.ID
	data := callback.Data

	dateIn, _ := userData[chatID]
	dateOut, _ := userData[chatID]

	if data == "date_in_incorrect" {
		bot.AnswerCallbackQuery(tgbotapi.NewCallback(callback.ID, "Попробуйте еще раз"))
		bot.DeleteMessage(tgbotapi.NewDeleteMessage(chatID, callback.Message.MessageID))
		StartSelectDateIn(bot, chatID, stateManager)
		return
	}

	if data == "date_in_correct" {
		bot.AnswerCallbackQuery(tgbotapi.NewCallback(callback.ID, "Укажите дату выезда"))
		bot.DeleteMessage(tgbotapi.NewDeleteMessage(chatID, callback.Message.MessageID))
		StartSelectDateOut(bot, chatID, dateIn, stateManager)
		return
	}

	if data == "date_out_incorrect" {
		bot.AnswerCallbackQuery(tgbotapi.NewCallback(callback.ID, "Попробуйте еще раз"))
		bot.DeleteMessage(tgbotapi.NewDeleteMessage(chatID, callback.Message.MessageID))
		StartSelectDateOut(bot, chatID, dateIn, stateManager)
		return
	}

	if data == "date_out_correct" {
		bot.AnswerCallbackQuery(tgbotapi.NewCallback(callback.ID, "Дата выезда указана"))
		bot.DeleteMessage(tgbotapi.NewDeleteMessage(chatID, callback.Message.MessageID))

		// Отправляем сообщение с итоговой информацией
		finalMsg := fmt.Sprintf(
			"<b>Дата заезда: </b>%s\n<b>Дата выезда: </b>%s\n\n<b>Все верно?</b>",
			dateIn.Format("02.01.2006"),
			dateOut.Format("02.01.2006"),
		)

		msg := tgbotapi.NewMessage(chatID, finalMsg)
		msg.ReplyMarkup = misc_utils.IsCorrectMarkup("city_info")
		msg.ParseMode = "HTML"
		bot.Send(msg)

		stateManager.SetState(chatID, states.IsInfoCorrect)
	}
}
