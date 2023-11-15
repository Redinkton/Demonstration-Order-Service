package repository

import (
	"database/sql"
	"main/internal/models"
)

type Repository struct {
	db *sql.DB
}

func (r *Repository) InsertOrder(order *models.Order){

}