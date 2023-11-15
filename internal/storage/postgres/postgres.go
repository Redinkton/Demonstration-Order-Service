package postgres

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const operation = "storage.postgres.New"
	
	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}
	
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}
	
	return &Storage{db: db}, nil
}

func (s *Storage) Close() {
	s.db.Close()
}