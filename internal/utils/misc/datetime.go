package misc

import (
	"fmt"
	"strings"
)

// Соответствие чисел и месяцев
var numbersToMonths = map[string]string{ // TODO - use format
	"01": "января",
	"02": "февраля",
	"03": "марта",
	"04": "апреля",
	"05": "мая",
	"06": "июня",
	"07": "июля",
	"08": "августа",
	"09": "сентября",
	"10": "октября",
	"11": "ноября",
	"12": "декабря",
}

// Преобразует строку с датой-временем в читаемый формат
func GetReadableDateTime(strDatetime string) string {
	parts := strings.Split(strDatetime, " ")
	if len(parts) < 2 {
		return strDatetime
	}

	date := parts[0]
	timeParts := strings.Split(parts[1], ".")[0]
	correctTimeParts := strings.Split(timeParts, ":")[:2]

	return fmt.Sprintf("%s, %s", strings.Join(correctTimeParts, ":"), GetReadableDate(date, "го"))
}

// Преобразует строку с датой в читаемый формат
func GetReadableDate(strDate string, ending string) string {
	parts := strings.Split(strDate, "-")
	if len(parts) < 3 {
		return strDate
	}

	year, month, day := parts[0], parts[1], parts[2]
	return fmt.Sprintf("%s-%s %s %s-го года", strings.TrimLeft(day, "0"), ending, numbersToMonths[month], year)
}
