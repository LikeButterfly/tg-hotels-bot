package rapidapi_models

type City struct {
	Name string // Короткое название города
	ID   int    // Идентификатор города (преобразованный из GaiaId)
}

type CitiesResponse struct {
	Q   string       `json:"q"`
	RC  string       `json:"rc"`
	Rid string       `json:"rid"`
	SR  []CityResult `json:"sr"`
}

type CityResult struct {
	EssId       EssId       `json:"essId"`
	GaiaId      string      `json:"gaiaId"`
	Index       string      `json:"index"`
	RegionNames RegionNames `json:"regionNames"`
}

type EssId struct {
	SourceId   string `json:"sourceId"`
	SourceName string `json:"sourceName"`
}

type RegionNames struct {
	DisplayName string `json:"displayName"`
	FullName    string `json:"fullName"`
	ShortName   string `json:"shortName"`
}
