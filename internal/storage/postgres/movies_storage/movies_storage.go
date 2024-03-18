package movies_storage

import "database/sql"

type MoviesStorage struct {
	db *sql.DB
}

func New(db *sql.DB) *MoviesStorage {
	return &MoviesStorage{db}
}
