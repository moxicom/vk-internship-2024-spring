package storage

import (
	"database/sql"

	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres/actors_storage"
)

type Actors interface {
	GetActors() ([]models.ActorFilm, error)
	GetActor(int) (models.ActorFilm, error)
	AddActor(models.Actor) (int, error)
	UpdateActor(int, models.Actor) error
	DeleteActor(int) error
}

type Movies interface {
}

type Repository struct {
	Actors
	// Movies
}

func New(db *sql.DB) *Repository {
	return &Repository{
		Actors: actors_storage.New(db),
	}
}
