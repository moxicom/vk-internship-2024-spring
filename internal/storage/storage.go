package storage

import (
	"database/sql"

	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres/actors_storage"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres/movies_storage"
)

type Actors interface {
	GetActors() ([]models.ActorFilms, error)
	GetActor(int) (models.ActorFilms, error)
	AddActor(models.Actor) (int, error)
	UpdateActor(int, models.Actor) error
	DeleteActor(int) error
}

type Movies interface {
	GetMovies(models.SortParams, models.SearchParams) ([]models.MovieActors, error)
	GetMovie(int) (models.MovieActors, error)
	AddMovie(models.Movie) (int, error)
	UpdateMovie(int, models.Movie) error
	DeleteMovie(int) error
}

type Repository struct {
	Actors
	Movies
}

func New(db *sql.DB) *Repository {
	return &Repository{
		Actors: actors_storage.New(db),
		Movies: movies_storage.New(db),
	}
}
