package movies_storage

import (
	"errors"
	"strconv"

	"github.com/moxicom/vk-internship-2024-spring/internal/models"
)

func (s *MoviesStorage) UpdateMovie(movieId int, movie models.Movie) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	queryParams := []interface{}{}
	updateQuery := "UPDATE movies SET"
	sqlArgIterator := 0

	if movie.Name != "" {
		sqlArgIterator++
		updateQuery += " name = $" + strconv.Itoa(sqlArgIterator) + ","
		queryParams = append(queryParams, movie.Name)
	}
	if movie.Description != "" {
		sqlArgIterator++
		updateQuery += " description = $" + strconv.Itoa(sqlArgIterator) + ","
		queryParams = append(queryParams, movie.Description)
	}
	if movie.Date != "" {
		sqlArgIterator++
		updateQuery += " date = $" + strconv.Itoa(sqlArgIterator) + ","
		queryParams = append(queryParams, movie.Date)
	}
	if movie.Rating != nil {
		sqlArgIterator++
		updateQuery += " rating = $" + strconv.Itoa(sqlArgIterator) + ","
		queryParams = append(queryParams, *movie.Rating)
	}

	// Remove the trailing comma if any fields are updated
	if len(queryParams) > 0 {
		updateQuery = updateQuery[:len(updateQuery)-1]
	} else {
		return errors.New("no fields to update")
	}

	sqlArgIterator++
	updateQuery += " WHERE movie_id = $" + strconv.Itoa(sqlArgIterator)
	queryParams = append(queryParams, movieId)

	stmt, err := tx.Prepare(updateQuery)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()
	// Execute the update movies
	_, err = stmt.Exec(queryParams...)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
