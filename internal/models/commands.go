package models

type CommandType string

const (
	HomeMenu  CommandType = "home_menu"
	Help      CommandType = "help"
	LowPrice  CommandType = "lowprice"
	HighPrice CommandType = "highprice"
	BestDeal  CommandType = "bestdeal"
	History   CommandType = "history"
	Favorites CommandType = "favorites"
)
