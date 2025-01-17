package models

type NearMarkets struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Location struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
}
