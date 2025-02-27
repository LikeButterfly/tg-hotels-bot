package config

type DefaultCommand struct {
	Command string `json:"command"`
	Desc    string `json:"description"`
}

var DefaultCommands = []DefaultCommand{
	{"start", "Запуск"},
	{"help", "Помощь"},
	{"lowprice", "Топ самых дешёвых отелей в городе"},
	// {"history", "История"},
}
