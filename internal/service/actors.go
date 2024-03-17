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

func (s *actorsService) GetActor(actorId int) (models.ActorFilm, error) {
	return s.storage.Actors.GetActor(actorId)
}

func (s *actorsService) AddActor(actor models.Actor) (int, error) {
	return s.storage.Actors.AddActor(actor)
}

func (s *actorsService) UpdateActor(actorId int, actor models.Actor) error {
	return s.storage.Actors.UpdateActor(actorId, actor)
}

func (s *actorsService) DeleteActor(actorId int) error {
	return s.storage.Actors.DeleteActor(actorId)
}
