package hotels_handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tg-hotels-bot/internal/config"
	"tg-hotels-bot/internal/models"
	"tg-hotels-bot/internal/states"
)

// Обрабатывает подтверждение информации о городе и датах
func ConfirmCityInfo(
	cfg *config.Config,
	bot *tgbotapi.BotAPI,
	callback *tgbotapi.CallbackQuery,
	stateManager *states.StateManager,
	usersData map[int64]*models.UserData,
) {
	chatID := callback.Message.Chat.ID

	if callback.Data == "city_info_correct" {
		if usersData[chatID] == nil {
			usersData[chatID] = &models.UserData{}
		}

		SendFirstHotel(cfg, bot, usersData[chatID], chatID)

	} else if callback.Data == "city_info_incorrect" {
		if _, err := bot.Send(tgbotapi.NewCallback(callback.ID, "Укажите информацию заново")); err != nil {
			log.Println("Ошибка ответа на callback: ", err)
		}
		editMsg := tgbotapi.NewEditMessageText(chatID, callback.Message.MessageID, "<b>Отправьте боту город для поиска</b>")
		editMsg.ParseMode = "HTML"
		bot.Send(editMsg)
		stateManager.SetState(chatID, states.WaitCityName)
	}
}
