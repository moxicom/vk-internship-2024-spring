package relations_storage

import "database/sql"

type RelationsStorage struct {
	db *sql.DB
}

func New(db *sql.DB) *RelationsStorage {
	return &RelationsStorage{db}
}
