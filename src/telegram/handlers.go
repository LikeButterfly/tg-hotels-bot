package telegram

import (
	"log"

	"context"
	"tg-hotels-bot/src/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) HandleCommands(updates tgbotapi.UpdatesChannel) (err error) {
	// Обрабатываем входящие сообщения
	for update := range updates {
		if update.Message == nil { // Игнорируем не-сообщения
			continue
		}

		switch update.Message.Text {
		case "/start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я твой ботяра")
			b.bot.Send(msg)

		case "/history":
			userID := update.Message.Chat.ID
			collection := database.GetHistoryCollection()

			var user struct {
				History map[string]map[string]any `bson:"history"`
			}
			err := collection.FindOne(context.TODO(), map[string]any{"_id": userID}).Decode(&user)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "История не найдена.")
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Println("Ошибка отправки сообщения /history:", err)
				}
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вот твоя история: TODO") // FIXME: отправка истории
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Println("Ошибка отправки истории:", err)
			}
		}
	}
	return nil
}
