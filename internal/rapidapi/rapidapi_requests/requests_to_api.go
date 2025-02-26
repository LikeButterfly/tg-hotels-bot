package rapidapi_requests

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"tg-hotels-bot/internal/config"
)

// RequestToAPI - GET-запрос к RapidAPI
func RequestToAPI(url string, queryParams map[string]string) (map[string]any, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = config.GetHeadersByCorrectRapidAPIKey()

	q := req.URL.Query()
	for key, value := range queryParams {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, errors.New("ошибка запроса к RapidAPI: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return map[string]any{"message": "No content"}, nil
	}
	if resp.StatusCode == http.StatusTooManyRequests {
		config.ChangeRapidAPIKey()
		return nil, errors.New("достигнут лимит запросов, ключ сменён")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("неожиданный код ответа: " + http.StatusText(resp.StatusCode))
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		return nil, errors.New("неожиданный content-type: " + contentType)
	}

	var responseJSON map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		return nil, errors.New("ошибка парсинга JSON: " + err.Error())
	}

	return responseJSON, nil
}

// RequestToAPIWithPayload - POST-запрос к RapidAPI
func RequestToAPIWithPayload(url string, payload map[string]any) (map[string]any, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	// Кодируем payload в JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.New("ошибка сериализации payload: " + err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	headers := config.GetHeadersWithPayloadByCorrectRapidAPIKey()
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, errors.New("ошибка запроса к RapidAPI: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return map[string]any{"message": "No content"}, nil
	}
	if resp.StatusCode == http.StatusTooManyRequests {
		config.ChangeRapidAPIKey()
		return nil, errors.New("достигнут лимит запросов, ключ сменён")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("неожиданный код ответа: " + http.StatusText(resp.StatusCode))
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		return nil, errors.New("неожиданный content-type: " + contentType)
	}

	var responseJSON map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&responseJSON); err != nil {
		return nil, errors.New("ошибка парсинга JSON: " + err.Error())
	}

	return responseJSON, nil
}
