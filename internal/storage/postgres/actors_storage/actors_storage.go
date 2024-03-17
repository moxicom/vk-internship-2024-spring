package actors_storage

import "database/sql"

type actorsStorage struct {
	db *sql.DB
}

func New(db *sql.DB) *actorsStorage {
	return &actorsStorage{db}
}
