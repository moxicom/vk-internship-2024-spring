package utils

import (
	"errors"
	"time"

	"github.com/moxicom/vk-internship-2024-spring/internal/models"
)

func ValidateMovie(movie models.Movie) error {
	// Validate time string if it is not empty
	if movie.Date != "" {
		_, err := time.Parse("2006-01-02", movie.Date)
		if err != nil {
			return errors.New("movie date must be in format YYYY-MM-DD")
		}
	}
	// Validate rating if it is not empty
	if movie.Rating != nil {
		if *movie.Rating < 0 || *movie.Rating > 10 {
			return errors.New("movie rating must be between 0 and 10")
		}
	}
	if len(movie.Name) > 150 {
		movieErr := errors.New("movie name max length is 150 characters")
		return movieErr
	}
	if len(movie.Description) > 1000 {
		movieErr := errors.New("movie description max length is 1000 characters")
		return movieErr
	}
	return nil
}
