package service

import (
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage"
)

type actorsService struct {
	storage *storage.Repository
}

func newActorsService(s *storage.Repository) *actorsService {
	return &actorsService{s}
}

func (s *actorsService) GetActors() ([]models.ActorFilm, error) {
	return s.storage.Actors.GetActors()
}
