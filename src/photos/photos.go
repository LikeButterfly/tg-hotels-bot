package photos

func GetPhotoByCommand(command string) string {
	photos := Photos

	if photo, exists := photos[command]; exists {
		return photo
	}

	// Если нет специфической картинки, берём дефолтную
	return photos["home_menu"]
}
