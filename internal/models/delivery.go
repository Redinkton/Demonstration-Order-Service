package models

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	City    string `json:"city"`
	Zip     int    `json:"zip"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}
