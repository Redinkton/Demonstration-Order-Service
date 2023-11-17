package storage

import (
	"database/sql"
	"log"
)

type PostgresRepository struct {
	db *sql.DB
}

func ConnectDB(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) Close() {
	err := r.db.Close()
	if err != nil {
		log.Println("Error closing database connection:", err)
	}
}
