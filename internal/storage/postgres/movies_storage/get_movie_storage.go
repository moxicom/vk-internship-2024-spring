package movies_storage

import (
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
)

func (s *MoviesStorage) GetMovie(movieId int) (models.MovieActors, error) {
	query := `SELECT movie_id, name, description, date, rating FROM movies WHERE movie_id = $1`
	row := s.db.QueryRow(query, movieId)
	movie := models.MovieActors{}
	err := row.Scan(&movie.ID, &movie.Name, &movie.Description, &movie.Date, &movie.Rating)
	if err != nil {
		return models.MovieActors{}, err
	}
	actors, err := s.GetMovieActors(movie.ID)
	if err != nil {
		return models.MovieActors{}, err
	}
	movie.Actors = actors
	return movie, nil
}
