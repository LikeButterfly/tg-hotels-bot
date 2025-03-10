package default_handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"tg-hotels-bot/internal/unsplash"
)

// Текст справки
var helpText = `<b>Поиск отелей:</b>
/lowprice - Поиск отелей по возрастанию цены
/highprice - Поиск отелей по убыванию цены
/bestdeal - Поиск отелей с условиями: диапазон цен, максимальная удаленность от центра

<b>Ваши отели:</b>
/history - История поиска с найденными отелями
/favorites - Отели, добавленные в избранное

<b>Другие команды:</b>
/start - Перезапуск бота
/help - Вывод данного сообщения

<b>Замечание:</b>
Поиск отелей в России временно недоступен

<b>Рекомендации:</b>
При возникновении ошибок:
1. Попробовать перезапустить бота, отправив /start
2. Если не помогает - попробовать снова через 2-5 минут
3. В других случаях можно написать разработчику (только если вы знаете его 😁)
`

func HandleHelp(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	photo := tgbotapi.NewPhoto(message.Chat.ID, tgbotapi.FileURL(unsplash.GetPhotoByCommand("help")))
	photo.Caption = helpText
	photo.ParseMode = "HTML"
	photo.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := bot.Send(photo); err != nil {
		// TODO обработать ошибку
		return
	}
}

func HandleHelpText(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	HandleHelp(bot, message)
}
