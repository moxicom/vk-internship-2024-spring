package movies_storage

func (s *moviesStorage) DeleteMovie(movieId int) error {
	// TODO: delete movie by id
	// TODO: do not forget to delete all relations with actors
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	// Delete movie reations with actors
	_, err = tx.Exec("DELETE FROM movie_actors WHERE movie_id=$1", movieId)
	if err != nil {
		tx.Rollback()
		return err
	}
	// Delete movie
	_, err = tx.Exec("DELETE FROM movies WHERE movie_id=$1", movieId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
