package models

// Содержит информацию об отеле
type HotelInfo struct {
	ID                 string
	Name               string
	Stars              int
	Address            string
	DistanceFromCenter float64
	TotalCost          float64
	CostByNight        float64
	Photo              string
	Latitude           float64
	Longitude          float64
}
