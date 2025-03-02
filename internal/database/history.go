package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"tg-hotels-bot/internal/config"
	mongodb_models "tg-hotels-bot/internal/models/mongodb"
	misc_utils "tg-hotels-bot/internal/utils/misc"
)

const (
	databaseName   = "Hotels"
	collectionName = "History"
)

// Возвращает коллекцию истории пользователя
func GetHistoryCollection(cfg *config.Config) (*mongo.Collection, error) {
	client, err := GetMongoClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к MongoDB: %w", err)
	}

	return client.Database(databaseName).Collection(collectionName), nil
}

// Добавляет команду в историю
func AddCommandToHistory(cfg *config.Config, command string, callTime time.Time, userID int64) error {
	collection, err := GetHistoryCollection(cfg)
	if err != nil {
		return err
	}

	var userDoc mongodb_models.UserHistoryDoc
	err = collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&userDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return createHistory(collection, command, callTime, userID)
		}
		log.Println("Ошибка при поиске пользователя в истории: ", err)
		return err
	}

	return addNewToHistory(&userDoc, collection, command, callTime, userID)
}

func addNewToHistory(userDoc *mongodb_models.UserHistoryDoc, collection *mongo.Collection, command string, callTime time.Time, userID int64) error {
	historyEntry := createHistoryEntry(command, callTime)
	if userDoc.History == nil {
		userDoc.History = make(map[string]mongodb_models.HistoryEntry)
	}
	key := callTime.Format(time.RFC3339)
	userDoc.History[key] = historyEntry

	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": userID},
		bson.M{"$set": bson.M{"history": userDoc.History}},
	)
	if err != nil {
		log.Println("Ошибка обновления истории: ", err)
	}
	return err
}

func createHistory(collection *mongo.Collection, command string, callTime time.Time, userID int64) error {
	historyEntry := createHistoryEntry(command, callTime)
	history := map[string]mongodb_models.HistoryEntry{
		callTime.Format(time.RFC3339): historyEntry,
	}

	_, err := collection.InsertOne(
		context.TODO(),
		bson.M{"_id": userID, "history": history},
	)
	if err != nil {
		log.Println("Ошибка при создании истории: ", err)
	}
	return err
}

func createHistoryEntry(command string, callTime time.Time) mongodb_models.HistoryEntry {
	return mongodb_models.HistoryEntry{
		Text: fmt.Sprintf("<b>Command</b> /%s called\nв %s",
			command, misc_utils.GetReadableDateTime(callTime.Format("2006-01-02 15:04:05"))),
		FoundHotels: []string{},
	}
}

// Добавляет город в историю пользователя
func AddCityToHistory(cfg *config.Config, city string, callTime time.Time, userID int64) error {
	collection, err := GetHistoryCollection(cfg)
	if err != nil {
		return err
	}

	var userDoc mongodb_models.UserHistoryDoc
	err = collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&userDoc)
	if err != nil {
		log.Println("Ошибка поиска пользователя в истории: ", err)
		return err
	}

	key := callTime.Format(time.RFC3339)
	if entry, exists := userDoc.History[key]; exists {
		entry.Text = fmt.Sprintf("Поиск в городе <b>%s</b>\n%s", city, entry.Text)
		userDoc.History[key] = entry
		_, err = collection.UpdateOne(
			context.TODO(),
			bson.M{"_id": userID},
			bson.M{"$set": bson.M{"history": userDoc.History}},
		)
		if err != nil {
			log.Println("Ошибка обновления истории с городом: ", err)
		}
	}

	return err
}
