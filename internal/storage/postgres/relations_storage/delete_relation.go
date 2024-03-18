package relations_storage

import "github.com/moxicom/vk-internship-2024-spring/internal/models"

func (s *RelationsStorage) DeleteRelation(rel models.RelationMoviesActors) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM movie_actors WHERE movie_id = $1 AND actor_id = $2", rel.MovieID, rel.ActorID)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
