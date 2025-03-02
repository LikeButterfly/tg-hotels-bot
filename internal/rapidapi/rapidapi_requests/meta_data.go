package rapidapi_requests

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"tg-hotels-bot/internal/config"
)

const metaDataURL = "https://hotels4.p.rapidapi.com/v2/get-meta-data" // FIXME тоже вынести бы по хорошему
const country = "US"

type CountryMetaData struct {
	SiteID           json.Number      `json:"siteId"`
	EAPID            json.Number      `json:"EAPID"`
	SupportedLocales []LocaleMetaData `json:"supportedLocales"`
}

type LocaleMetaData struct {
	LocaleIdentifier   string      `json:"localeIdentifier"`
	LanguageIdentifier json.Number `json:"languageIdentifier"`
}

// Запрашивает мета-данные и устанавливает переменные окружения
func FetchMetaData() error {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", metaDataURL, nil)
	if err != nil {
		log.Println("Ошибка при создании запроса:", err)
		return err
	}
	req.Header = config.GetHeadersByCorrectRapidAPIKey()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		log.Println("Ошибка запроса к RapidAPI:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		config.ChangeRapidAPIKey()
		return fmt.Errorf("rate limit exceeded")
	}

	// Декодируем ответ в типизированную структуру.
	var metaData map[string]CountryMetaData
	if err := json.NewDecoder(resp.Body).Decode(&metaData); err != nil {
		log.Println("Ошибка декодирования JSON:", err)
		return err
	}

	countryData, exists := metaData[country]
	if !exists {
		log.Println("Ошибка: данные о стране не найдены")
		return fmt.Errorf("country data not found")
	}

	setEnvRapidAPIVars(countryData)
	return nil
}

// Устанавливает переменные окружения
func setEnvRapidAPIVars(data CountryMetaData) {
	os.Setenv("SITE_ID", data.SiteID.String())
	os.Setenv("EAPID", data.EAPID.String())
	if len(data.SupportedLocales) > 0 {
		locale := data.SupportedLocales[0]
		os.Setenv("LOCALE_IDENTIFIER", locale.LocaleIdentifier)
		os.Setenv("LANGUAGE_IDENTIFIER", locale.LanguageIdentifier.String())
	}
}
