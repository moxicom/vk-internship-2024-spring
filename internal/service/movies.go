package service

import (
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage"
)

type moviesService struct {
	storage *storage.Repository
}

func newMoviesService(s *storage.Repository) *moviesService {
	return &moviesService{s}
}

func (s *moviesService) GetMovies(sort models.SortParams, search models.SearchParams) ([]models.MovieActors, error) {
	return s.storage.Movies.GetMovies(sort, search)
}

func (s *moviesService) GetMovie(movieId int) (models.MovieActors, error) {
	return s.storage.Movies.GetMovie(movieId)
}

func (s *moviesService) AddMovie(movie models.Movie) (int, error) {
	return s.storage.Movies.AddMovie(movie)
}

func (s *moviesService) UpdateMovie(movieId int, movie models.Movie) error {
	return s.storage.Movies.UpdateMovie(movieId, movie)
}

func (s *moviesService) DeleteMovie(movieId int) error {
	return s.storage.Movies.DeleteMovie(movieId)
}
