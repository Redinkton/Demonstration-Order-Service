package models

type Items struct {
	ChrtId      int     `json:"chrt_id"`
	TrackNumber Order   `json:"track_number"`
	Price       float32 `json:"price"`
	Rid         string  `json:"rid"`
	Name        string  `json:"name"`
	Sale        int     `json:"sale"`
	Size        int     `json:"size"`
	TotalPrice  float32 `json:"total_price"`
	NmId        int     `json:"nm_id"`
	Brand       string  `json:"brand"`
	Status      int     `json:"status"`
}
