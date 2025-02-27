package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"tg-hotels-bot/internal/config"
	misc_utils "tg-hotels-bot/internal/utils/misc"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseName   = "Hotels"
	collectionName = "History"
)

// Возвращает коллекцию истории пользователя
func GetHistoryCollection(cfg *config.Config) (*mongo.Collection, error) {
	client, err := GetMongoClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("Ошибка подключения к MongoDB: ", err)
	}

	return client.Database(databaseName).Collection(collectionName), nil
}

// Добавляет команду в историю
func AddCommandToHistory(cfg *config.Config, command string, callTime time.Time, userID int64) error {
	collection, err := GetHistoryCollection(cfg)
	if err != nil {
		return err
	}

	// Проверяем, есть ли пользователь в базе
	var user struct {
		History map[string]map[string]any `bson:"history"`
	}
	err = collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return createHistory(collection, command, callTime, userID)
		}
		log.Println("Ошибка при поиске пользователя в истории:", err)
		return err
	}

	return addNewToHistory(user.History, collection, command, callTime, userID)
}

func addNewToHistory(userHistory map[string]map[string]any, collection *mongo.Collection, command string, callTime time.Time, userID int64) error {
	historyEntry := createHistoryDict(command, callTime)
	userHistory[callTime.Format(time.RFC3339)] = historyEntry

	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": userID}, bson.M{"$set": bson.M{"history": userHistory}})
	if err != nil {
		log.Println("Ошибка обновления истории:", err)
	}
	return err
}

func createHistory(collection *mongo.Collection, command string, callTime time.Time, userID int64) error {
	historyEntry := createHistoryDict(command, callTime)
	history := map[string]map[string]any{callTime.Format(time.RFC3339): historyEntry}

	_, err := collection.InsertOne(context.TODO(), bson.M{"_id": userID, "history": history})
	if err != nil {
		log.Println("Ошибка при создании истории:", err)
	}
	return err
}

func createHistoryDict(command string, callTime time.Time) map[string]any {
	return map[string]any{
		"text":         fmt.Sprintf("<b>Command</b> /%s called\nв %s", command, misc_utils.GetReadableDateTime(callTime.Format("2006-01-02 15:04:05"))),
		"found_hotels": []string{},
	}
}

// Добавляет город в историю пользователя
func AddCityToHistory(cfg *config.Config, city string, callTime time.Time, userID int64) error {
	collection, err := GetHistoryCollection(cfg)
	if err != nil {
		return err
	}

	var user struct {
		History map[string]map[string]any `bson:"history"`
	}
	err = collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		log.Println("Ошибка поиска пользователя в истории:", err)
		return err
	}

	callTimeStr := callTime.Format(time.RFC3339)
	if historyPage, exists := user.History[callTimeStr]; exists {
		historyPage["text"] = fmt.Sprintf("Поиск в городе <b>%s</b>\n%s", city, historyPage["text"])
		_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": userID}, bson.M{"$set": bson.M{"history": user.History}})
		if err != nil {
			log.Println("Ошибка обновления истории с городом:", err)
		}
	}

	return err
}
