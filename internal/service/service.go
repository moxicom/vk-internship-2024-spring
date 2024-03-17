package service

import (
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Actors interface {
	GetActors() ([]models.ActorFilm, error)
	GetActor(int) (models.ActorFilm, error)
	AddActor(models.Actor) (int, error)
	UpdateActor(int, models.Actor) error
	DeleteActor(int) error
}

type Service struct {
	Actors
}

func New(s *storage.Repository) *Service {
	return &Service{
		Actors: newActorsService(s),
	}
}
