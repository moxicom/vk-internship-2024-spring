package movies_storage

import "database/sql"

type moviesStorage struct {
	db *sql.DB
}

func New(db *sql.DB) *moviesStorage {
	return &moviesStorage{db}
}
