package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"tg-hotels-bot/internal/config"
)

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

func (b *Bot) SetDefaultCommands(cfg *config.Config) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/setMyCommands", b.BotAPI.Token) // FIXME??

	var commands []BotCommand
	for cmd, desc := range cfg.DefaultCommands {
		commands = append(commands, BotCommand{Command: cmd, Description: desc})
	}

	payload, err := json.Marshal(map[string]any{
		"commands": commands,
	})
	if err != nil {
		log.Println("Ошибка сериализации команд:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Println("Ошибка запроса к Telegram API:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Ошибка установки команд: %d\n", resp.StatusCode)
		return
	}

	log.Println("Команды успешно установлены")
}
