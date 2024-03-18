package movies_storage

import (
	"fmt"

	"github.com/moxicom/vk-internship-2024-spring/internal/models"
)

func (s *MoviesStorage) GetMovies(sortParams models.SortParams, searchParams models.SearchParams) ([]models.MovieActors, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	query := `SELECT m.movie_id, m.name, m.description, m.date, m.rating
	              FROM movies m 
	              LEFT JOIN movie_actors ma ON m.movie_id = ma.movie_id
				  LEFT JOIN actors a ON ma.actor_id = a.actor_id
	              WHERE 1=1`
	args := []interface{}{}

	sqlArgIterator := 0
	if searchParams.MovieName != "" {
		sqlArgIterator++
		query += " AND LOWER(m.name) LIKE LOWER($" + fmt.Sprintf("%v", sqlArgIterator) + ")"
		args = append(args, "%"+searchParams.MovieName+"%")
	}

	if searchParams.ActorName != "" {
		sqlArgIterator++
		query += " AND LOWER(a.name) LIKE LOWER($" + fmt.Sprintf("%v", sqlArgIterator) + ")"
		args = append(args, "%"+searchParams.ActorName+"%")
	}

	if sortParams.Sort != "" && sortParams.Order != "" {
		query += " ORDER BY " + sortParams.Sort + " " + sortParams.Order
	} else {
		query += " ORDER BY m.rating DESC"
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	rows, err := stmt.Query(args...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stmt.Close()
	defer rows.Close()

	var movies []models.MovieActors
	for rows.Next() {
		var movie models.MovieActors
		if err := rows.Scan(&movie.ID, &movie.Name, &movie.Description, &movie.Date, &movie.Rating); err != nil {
			return nil, err
		}
		actors, err := s.GetMovieActors(movie.ID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		movie.Actors = actors
		movies = append(movies, movie)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return movies, nil
}

func (s *MoviesStorage) GetMovieActors(movieId int) ([]int, error) {
	actorsRows, err := s.db.Query("SELECT actor_id from movie_actors WHERE movie_id=$1", movieId)
	if err != nil {
		return []int{}, err
	}
	defer actorsRows.Close()

	var actors []int
	for actorsRows.Next() {
		var actor_id int
		err := actorsRows.Scan(&actor_id)
		if err != nil {
			return []int{}, err
		}
		actors = append(actors, actor_id)
	}
	if err := actorsRows.Err(); err != nil {
		return []int{}, err
	}
	return actors, nil
}
