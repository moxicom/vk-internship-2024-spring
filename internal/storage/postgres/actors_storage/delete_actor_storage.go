package actors_storage

func (s *actorsStorage) DeleteActor(actorId int) error {
	// Delete actor
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM movie_actors WHERE actor_id=$1", actorId)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec("DELETE FROM actors WHERE actor_id=$1", actorId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
