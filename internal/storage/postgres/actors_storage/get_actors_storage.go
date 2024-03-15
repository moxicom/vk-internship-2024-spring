package actors_storage

import (
	"database/sql"

	"github.com/moxicom/vk-internship-2024-spring/internal/models"
)

type actorsStorage struct {
	db *sql.DB
}

func New(db *sql.DB) *actorsStorage {
	return &actorsStorage{db}
}

func (s *actorsStorage) GetActors() ([]models.ActorFilm, error) {
	actorsRows, err := s.db.Query("SELECT actor_id, name, isMale, date_of_birth FROM actors")
	if err != nil {
		return []models.ActorFilm{}, err
	}
	defer actorsRows.Close()

	var actors []models.ActorFilm
	for actorsRows.Next() {
		var actor models.ActorFilm
		// Scan actor
		err := actorsRows.Scan(&actor.ID, &actor.Name, &actor.IsMale, &actor.BirthDay)
		if err != nil {
			return []models.ActorFilm{}, err
		}
		// Get actors fims
		movies, err := s.getActorFilmsIDs(int(actor.ID))
		if err != nil {
			return []models.ActorFilm{}, nil
		}
		actor.Movies = movies
		actors = append(actors, actor)
	}
	if err := actorsRows.Err(); err != nil {
		return []models.ActorFilm{}, err
	}

	return actors, nil
}

func (s *actorsStorage) getActorFilmsIDs(actorID int) ([]int, error) {
	// Get all filst of user
	movieRows, err := s.db.Query("SELECT movie_id from movie_actors WHERE actor_id=$1", actorID)
	if err != nil {
		return []int{}, err
	}
	defer movieRows.Close()

	var movies []int
	for movieRows.Next() {
		var movieId int
		err := movieRows.Scan(&movieId)
		if err != nil {
			return []int{}, err
		}
		movies = append(movies, movieId)
	}
	if err := movieRows.Err(); err != nil {
		return []int{}, err
	}
	return movies, nil
}

// func (s *actorsStorage)
