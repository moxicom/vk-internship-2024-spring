package movies_storage

import "github.com/moxicom/vk-internship-2024-spring/internal/models"

func (s *moviesStorage) GetMovies(sort models.SortParams) ([]models.MovieActors, error) {
	return []models.MovieActors{}, nil
}
