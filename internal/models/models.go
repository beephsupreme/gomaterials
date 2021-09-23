package models

// Data holds records from data.csv
type Data struct {
	PartNum string
	OnHand  float64
	OnOrder float64
	ReOrder float64
}
