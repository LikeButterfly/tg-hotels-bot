package config

import "os"

func GetCorrectRapidAPIKey() string {
	if os.Getenv("REQUESTS_LIMIT_REACHED") == "False" {
		return os.Getenv("RAPID_API_KEY")
	}
	return os.Getenv("ANOTHER_RAPID_API_KEY")
}

func ChangeRapidAPIKey() {
	if os.Getenv("REQUESTS_LIMIT_REACHED") == "False" {
		os.Setenv("REQUESTS_LIMIT_REACHED", "True")
	} else {
		os.Setenv("REQUESTS_LIMIT_REACHED", "False")
	}
}

func GetHeadersByCorrectRapidAPIKey() map[string]string {
	return map[string]string{
		"X-RapidAPI-Key":  GetCorrectRapidAPIKey(),
		"X-RapidAPI-Host": "hotels4.p.rapidapi.com",
	}
}

func GetHeadersWithPayloadByCorrectRapidAPIKey() map[string]string {
	return map[string]string{
		"content-type":    "application/json",
		"X-RapidAPI-Key":  GetCorrectRapidAPIKey(),
		"X-RapidAPI-Host": "hotels4.p.rapidapi.com",
	}
}
