package relations_storage

import "github.com/moxicom/vk-internship-2024-spring/internal/models"

func (s *RelationsStorage) AddRelation(rel models.RelationMoviesActors) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO movie_actors (movie_id, actor_id) VALUES ($1, $2)", rel.MovieID, rel.ActorID)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
