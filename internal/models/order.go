package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderUid          uuid.UUID `json:"-"`
	TrackNumber       string    `json:"track_number"`
	Entry             Entry     `json:"entry"`
	Delivery          Delivery  `json:"delivery"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        uuid.UUID `json:"-"`
	DeliveryService   string    `json:"delivery_service"`
	ShardKey          int       `json:"shardkey"`
	SmId              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          int       `json:"oof_shard"`
}
