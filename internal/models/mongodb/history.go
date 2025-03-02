package mongodb_models

// Запись в истории
type HistoryEntry struct {
	Text        string   `bson:"text"`
	FoundHotels []string `bson:"found_hotels"`
}

// Документ пользователя с историей команд
type UserHistoryDoc struct {
	ID      int64                   `bson:"_id"`
	History map[string]HistoryEntry `bson:"history"`
}
