package handlers

import (
	"log"
	"strings"

	"tg-hotels-bot/src/states"
	"tg-hotels-bot/src/telegram/handlers/action_handlers"
	"tg-hotels-bot/src/telegram/handlers/default_handlers"
	"tg-hotels-bot/src/telegram/handlers/hotels_handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// HandleCommands обрабатывает входящие сообщения и команды
func HandleCommands(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel, stateManager *states.StateManager, userData map[int64]map[string]string) {
	for update := range updates {
		// === ОБРАБОТКА CALLBACK-ЗАПРОСОВ (Inline-кнопки) ===

		if update.CallbackQuery != nil {
			chatID := update.CallbackQuery.Message.Chat.ID
			callbackData := update.CallbackQuery.Data

			log.Printf("Получен callback-запрос от %d: %s", chatID, callbackData)

			// Если пользователь выбирает город
			if strings.HasPrefix(callbackData, "search_in_city") {
				hotels_handlers.SetCityID(bot, update.CallbackQuery, stateManager, userData)
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

		// Команды поиска отелей (пока только lowprice)
		case "/lowprice":
			action_handlers.DefineState(bot, update.Message, stateManager, userData)

		default:
			if strings.HasPrefix(text, "/") {
				log.Printf("Неизвестная команда: %s", text)
				msg := tgbotapi.NewMessage(chatID, "Неизвестная команда. Используйте /help для списка доступных команд.")
				bot.Send(msg)
			} else {
				state, exists := stateManager.GetState(chatID)
				if exists {
					switch state {
					case states.WaitCityName:
						hotels_handlers.GetCitiesByName(bot, update.Message, stateManager)
					default:
						log.Printf("Необработанное сообщение: %s", text)
					}
				}
			}
		}
	}
}
