package default_handlers

import (
	"fmt"
	"tg-hotels-bot/src/utils/work_with_messages"

	"tg-hotels-bot/src/states"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// FinishWithError отправляет сообщение об ошибке и завершает сценарий
func FinishWithError(bot *tgbotapi.BotAPI, chatID int64, errorText string, stateManager *states.StateManager, toDelete ...int) {
	// Отправляем сообщение об ошибке
	msg := tgbotapi.NewMessage(chatID, createErrorMessage(errorText))
	msg.ParseMode = "HTML"
	bot.Send(msg)

	// Удаляем сообщения ожидания (если есть)
	for _, msgID := range toDelete {
		work_with_messages.DeleteMessage(bot, chatID, msgID)
	}

	// Возвращаем в главное меню
	HandleStart(bot, &tgbotapi.Message{Chat: &tgbotapi.Chat{ID: chatID}}, stateManager)
}

// createErrorMessage создаёт текст ошибки
func createErrorMessage(errorText string) string {
	template := "⚠️ <b>%s</b>"

	switch errorText {
	case "cities_not_found":
		return fmt.Sprintf(template, "Городов с таким названием не найдено")
	case "hotels_not_found":
		return fmt.Sprintf(template, "Отелей с заданными условиями не найдено")
	case "favorites_empty":
		return fmt.Sprintf(template, "Список избранного пуст")
	case "history_empty":
		return fmt.Sprintf(template, "История пуста")
	case "empty":
		return fmt.Sprintf(template, "Произошла ошибка при получении информации о городах. Попробуйте еще раз")
	case "timeout":
		return fmt.Sprintf(template, "Произошла ошибка на сервере. Попробуйте еще раз")
	case "page_index":
		return fmt.Sprintf(template, "Найденные отели закончились")
	case "bad_result":
		return fmt.Sprintf(template, "Возникла ошибка при получении информации. Попробуйте еще раз")
	default:
		return fmt.Sprintf(template, "Неизвестная ошибка")
	}
}
