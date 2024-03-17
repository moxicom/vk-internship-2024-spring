package service

import (
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Actors interface {
	GetActors() ([]models.ActorFilms, error)
	GetActor(int) (models.ActorFilms, error)
	AddActor(models.Actor) (int, error)
	UpdateActor(int, models.Actor) error
	DeleteActor(int) error
}

type Movies interface {
	GetMovies(models.SortParams) ([]models.MovieActors, error)
	GetMovie(int) (models.MovieActors, error)
	AddMovie(models.Movie) (int, error)
	UpdateMovie(int, models.Movie) error
	DeleteMovie(int) error
}

type Service struct {
	Actors
	Movies
}

func New(s *storage.Repository) *Service {
	return &Service{
		Actors: newActorsService(s),
		Movies: newMoviesService(s),
	}
}
