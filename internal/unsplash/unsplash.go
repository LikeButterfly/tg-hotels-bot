package unsplash

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"tg-hotels-bot/internal/models"
)

// Фолбэк картинки, если Unsplash не вернул результат
var Photos = map[models.CommandType]string{
	models.HomeMenu:  "https://i.pinimg.com/736x/5a/f3/3f/5af33f8cf06f6061a36191c3f88c96e9.jpg",
	models.Help:      "https://i.pinimg.com/originals/af/be/b3/afbeb32ecbec0f95d1d37b75d8f428b5.jpg",
	models.LowPrice:  "https://i.pinimg.com/originals/e1/7d/81/e17d812bff962652884b2f6c366e602d.jpg",
	models.HighPrice: "https://i.pinimg.com/originals/f1/a5/d9/f1a5d9597183f133f290068f07eedb94.jpg",
	models.BestDeal:  "https://q-xx.bstatic.com/xdata/images/hotel/max1024x768/40044571.jpg?k=0345ddebf079962f78134757597863b0dc244edd52d278640ad364e83a43ebfe&o=",
	models.History:   "https://i.pinimg.com/originals/32/ca/ed/32caede5cc406cf8538e7377022b0cd2.jpg",
	models.Favorites: "https://avatars.mds.yandex.net/get-altay/4365309/2a0000017985c68c274649deba476f9504eb/XXL",
}

// Для каждой команды задаем адекватный поисковый запрос на Unsplash
var searchQueries = map[models.CommandType]string{
	models.HomeMenu:  "modern hotel lobby",
	models.Help:      "help desk in modern hotel",
	models.LowPrice:  "affordable hotel room interior",
	models.HighPrice: "luxury hotel suite interior",
	models.BestDeal:  "cozy boutique hotel exterior",
	models.History:   "historic hotel building exterior",
	models.Favorites: "popular hotel exterior",
}

// Структура ответа от Unsplash.
type UnsplashResponse struct {
	URLs struct {
		Regular string `json:"regular"`
	} `json:"urls"`
}

const UnsplashURL = "https://api.unsplash.com/photos/random"

// Возвращает URL фотографии, полученной с Unsplash по указанной команде.
// Если Unsplash не отвечает, возвращается базовая картинка из Photos
func GetPhotoByCommand(cmd models.CommandType) string {
	query, exists := searchQueries[cmd]
	if !exists {
		return Photos[models.HomeMenu]
	}
	UnsplashAccessKey := os.Getenv("UNSPLASH_ACCESS_KEY")
	if UnsplashAccessKey == "" {
		return Photos[cmd]
	}

	photo, err := getRandomPhoto(query, UnsplashAccessKey)
	if err != nil || photo == "" {
		return Photos[cmd]
	}
	return photo
}

func getRandomPhoto(query, accessKey string) (string, error) {
	req, err := http.NewRequest("GET", UnsplashURL, nil)
	if err != nil {
		log.Println("Error in getRandomPhoto: ", err)
		return "", err
	}
	q := req.URL.Query()
	q.Add("query", query)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", "Client-ID "+accessKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unsplash API error: %s", resp.Status)
	}

	var result UnsplashResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.URLs.Regular, nil
}
