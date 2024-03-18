package movies_storage

func (s *MoviesStorage) DeleteMovie(movieId int) error {
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
