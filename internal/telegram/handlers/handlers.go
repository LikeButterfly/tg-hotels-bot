package handlers

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tg-hotels-bot/internal/config"
	"tg-hotels-bot/internal/models"
	"tg-hotels-bot/internal/states"
	"tg-hotels-bot/internal/telegram/handlers/action_handlers"
	"tg-hotels-bot/internal/telegram/handlers/default_handlers"
	"tg-hotels-bot/internal/telegram/handlers/hotels_handlers"
)

// Обрабатывает входящие сообщения и команды
func HandleCommands(
	cfg *config.Config,
	bot *tgbotapi.BotAPI,
	updates tgbotapi.UpdatesChannel,
	stateManager *states.StateManager,
	usersData map[int64]*models.UserData,
) {
	for update := range updates {
		// === ОБРАБОТКА CALLBACK-ЗАПРОСОВ (Inline-кнопки) ===

		if update.CallbackQuery != nil {
			chatID := update.CallbackQuery.Message.Chat.ID
			callbackData := update.CallbackQuery.Data

			log.Printf("Получен callback-запрос от %d: %s", chatID, callbackData)

			// Если пользователь выбирает город
			if strings.HasPrefix(callbackData, "search_in_city") {
				hotels_handlers.SetCityID(bot, update.CallbackQuery, stateManager, usersData)
				continue
			}

			// Подтверждение пользователем выбора города и дат
			if strings.HasPrefix(callbackData, "city_info") {
				hotels_handlers.ConfirmCityInfo(cfg, bot, update.CallbackQuery, stateManager, usersData)
				continue
			}
		}

		// === ОБРАБОТКА ТЕКСТОВЫХ СООБЩЕНИЙ ===

		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		text := update.Message.Text

		log.Printf("Получена команда от %d: %s", chatID, text)

		switch text {
		// Базовые команды
		case "/start":
			default_handlers.HandleStart(bot, update.Message, stateManager)
		case "Главное меню":
			default_handlers.HandleMainMenu(bot, update.Message, stateManager)
		case "/help", "Справка":
			default_handlers.HandleHelp(bot, update.Message)

		case "/lowprice", "Недорогие отели":
			action_handlers.ShowLowprice(cfg, bot, update.Message, stateManager, usersData)

		case "Показать еще":
			hotels_handlers.ShowMoreHotels(cfg, bot, usersData[chatID], chatID) // FIXME

		default:
			if strings.HasPrefix(text, "/") {
				log.Printf("Неизвестная команда: %s", text)
				msg := tgbotapi.NewMessage(chatID, "Неизвестная команда. Используйте /help для списка доступных команд.")
				if _, err := bot.Send(msg); err != nil {
					// TODO обработать ошибку
					return
				}
			} else {
				state, exists := stateManager.GetState(chatID)
				if exists {
					switch state {
					case states.WaitCityName:
						hotels_handlers.GetCitiesByName(cfg, bot, update.Message, stateManager)

					// Ожидаем ввод даты заезда
					case states.SelectDateIn:
						hotels_handlers.SelectDateIn(bot, update.Message, stateManager, usersData)

					// Ожидаем ввод даты выезда
					case states.SelectDateOut:
						hotels_handlers.SelectDateOut(bot, update.Message, stateManager, usersData)

					default:
						log.Printf("Необработанное сообщение: %s", text)
					}
				}
			}
		}
	}
}
