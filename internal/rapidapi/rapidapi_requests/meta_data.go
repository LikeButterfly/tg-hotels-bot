package rapidapi_requests

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"tg-hotels-bot/internal/config"
	"time"
)

const metaDataURL = "https://hotels4.p.rapidapi.com/v2/get-meta-data" // FIXME тоже вынести бы по хорошему
const country = "US"

// Запрашивает мета-данные и устанавливает переменные окружения
func FetchMetaData() error {
	client := &http.Client{
		Timeout: 10 * time.Second, // Таймаут запроса
	}

	req, err := http.NewRequest("GET", metaDataURL, nil)
	if err != nil {
		log.Println("Ошибка при создании запроса:", err)
		return err
	}

	// Устанавливаем заголовки RapidAPI
	req.Header = config.GetHeadersByCorrectRapidAPIKey()

	// Отправляем запрос
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		log.Println("Ошибка запроса к RapidAPI:", err)
		return err
	}
	defer resp.Body.Close()

	// Если API вернул 429 (лимит запросов) — меняем ключ
	if resp.StatusCode == http.StatusTooManyRequests {
		config.ChangeRapidAPIKey()
		return err
	}

	// Декодируем JSON-ответ
	var responseJSON map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		log.Println("Ошибка декодирования JSON:", err)
		return err
	}

	// Извлекаем данные по стране
	countryData, exists := responseJSON[country].(map[string]any)
	if !exists {
		log.Println("Ошибка: данные о стране не найдены")
		return err
	}

	// Устанавливаем переменные окружения
	setEnvRapidAPIVars(countryData)

	return nil
}

// Устанавливает переменные окружения
func setEnvRapidAPIVars(countryData map[string]any) {
	os.Setenv("SITE_ID", toString(countryData["siteId"]))
	os.Setenv("EAPID", toString(countryData["EAPID"]))

	if supportedLocales, ok := countryData["supportedLocales"].([]any); ok && len(supportedLocales) > 0 {
		if localeData, ok := supportedLocales[0].(map[string]any); ok {
			os.Setenv("LOCALE_IDENTIFIER", toString(localeData["localeIdentifier"]))
			os.Setenv("LANGUAGE_IDENTIFIER", toString(localeData["languageIdentifier"]))
		}
	}
}

// Конвертирует значение в строку
func toString(value any) string { // FIXME избавиться
	if str, ok := value.(string); ok {
		return str
	}
	if num, ok := value.(float64); ok {
		return fmt.Sprintf("%.0f", num) // Преобразование float в int -> string
	}
	return ""
}
