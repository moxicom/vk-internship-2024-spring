package utils

import (
	"errors"
	"strconv"

	"github.com/moxicom/vk-internship-2024-spring/internal/models"
)

func ValidateRelationParams(rel models.RelationMoviesActors) error {
	if rel.ActorID == "" && rel.MovieID == "" {
		return errors.New("invalid relation params. both are empty")
	}
	if rel.ActorID != "" {
		_, err := strconv.Atoi(rel.ActorID)
		if err != nil {
			return errors.New("invalid actor id")
		}
	}
	if rel.MovieID != "" {
		_, err := strconv.Atoi(rel.MovieID)
		if err != nil {
			return errors.New("invalid movie id")
		}
	}
	return nil
}
