package storage

import (
	"database/sql"
	"fmt"
	"main/internal/models"
)

type OrderRepository struct {
	db *sql.DB
}

func (o *OrderRepository) Create(order models.Order) error {
	query := "INSERT INTO order (order_uid, track_number, entry, delivery, locale, internal_signature, customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := o.db.Exec(query, order.OrderUid, order.TrackNumber, order.Entry, order.Delivery, order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService, order.ShardKey, order.SmId, order.DateCreated, order.OofShard)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderRepository) GetAll() ([]models.Order, error) {
	rows, err := o.db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.OrderUid, &order.TrackNumber, &order.Entry, &order.Delivery, &order.Locale, &order.InternalSignature, &order.CustomerId, &order.DeliveryService, &order.ShardKey, &order.SmId, &order.DateCreated, &order.OofShard)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (o *OrderRepository) Get(uid string) (models.Order, error) {
	var order models.Order
	err := o.db.QueryRow("SELECT * FROM orders WHERE order_uid = ?", uid).Scan(&order.OrderUid)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Order{}, fmt.Errorf("order not found")
		}
		return models.Order{}, err
	}
	return order, nil
}