package models

import "github.com/google/uuid"

type Items struct {
	ChrtId      int       `json:"chrt_id"`
	TrackNumber Order     `json:"track_number"`
	Price       float32   `json:"price"`
	RID         uuid.UUID `json:"rid"`
	Name        string    `json:"name"`
	Sale        int       `json:"sale"`
	Size        string    `json:"size"`
	TotalPrice  float32   `json:"total_price"`
	NmId        int       `json:"nm_id"`
	Brand       string    `json:"brand"`
	Status      int       `json:"status"`
}
