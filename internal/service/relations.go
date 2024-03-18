package service

import (
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage"
)

type relationsService struct {
	storage *storage.Repository
}

func newRelationsService(s *storage.Repository) *relationsService {
	return &relationsService{s}
}

func (s *relationsService) AddRelation(rel models.RelationMoviesActors) error {
	// TODO: ADD ADD RELATIONS
	return nil
	// return s.storage.Relations.AddRelation(rel)
}

func (s *relationsService) DeleteRelation(rel models.RelationMoviesActors) error {
	// TODO: ADD DELETE RELATIONS
	return nil
	// return s.storage.Relations.DeleteRelation(rel)
}
