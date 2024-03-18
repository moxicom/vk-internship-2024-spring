package storage

import (
	"database/sql"

	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres/actors_storage"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres/movies_storage"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres/relations_storage"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres/users_storage"
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

type Relations interface {
	AddRelation(models.RelationMoviesActors) error
	DeleteRelation(models.RelationMoviesActors) error
}

type Users interface {
	CheckUser(string, string, bool) (bool, error)
}

type Repository struct {
	Actors
	Movies
	Relations
	Users
}

func New(db *sql.DB) *Repository {
	return &Repository{
		Actors:    actors_storage.New(db),
		Movies:    movies_storage.New(db),
		Relations: relations_storage.New(db),
		Users:     users_storage.New(db),
	}
}
