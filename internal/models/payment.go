package models

import "github.com/google/uuid"

type Payment struct {
	Transaction  Order     `json:"transaction"`
	RequestId    uuid.UUID `json:"request_id"`
	Currency     Currency  `json:"currency"`
	Provider     Provider  `json:"provider"`
	Amount       int       `json:"amount"`
	PaymentDt    int       `json:"payment_dt"`
	Bank         string    `json:"bank"`
	DeliveryCost int       `json:"delivery_cost"`
	GoodsTotal   int       `json:"goods_total"`
	CustomFee    int       `json:"custom_fee"`
}
