package service

import (
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage"
)

type Actors interface {
	GetActors() ([]models.ActorFilm, error)
	AddActor(models.Actor) (int, error)
}

type Service struct {
	Actors
}

func New(s *storage.Repository) *Service {
	return &Service{
		Actors: newActorsService(s),
	}
}
