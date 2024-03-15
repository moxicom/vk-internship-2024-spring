package storage

import (
	"database/sql"

	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres"
)

type Actors interface {
	GetActors() ([]models.ActorFilm, error)
}

type Movies interface {
}

type Repository struct {
	Actors
	// Movies
}

func New(db *sql.DB) *Repository {
	return &Repository{
		Actors: postgres.NewActorsStorage(db),
	}
}
