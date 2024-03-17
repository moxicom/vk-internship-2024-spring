package movies_storage

import "github.com/moxicom/vk-internship-2024-spring/internal/models"

func (s *moviesStorage) AddMovie(movie models.Movie) (int, error) {
	var id int
	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}

	err = tx.QueryRow(
		"INSERT INTO movies(name, description, date, rating) VALUES($1, $2, $3, $4) RETURNING movie_id",
		movie.Name,
		movie.Description,
		movie.Date,
		movie.Rating).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
