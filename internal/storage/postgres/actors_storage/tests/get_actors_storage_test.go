package actors_storage_test

import (
	"errors"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres/actors_storage"
	"github.com/stretchr/testify/assert"
)

func TestActorsStorage_GetActors(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := actors_storage.New(db)

	tests := []struct {
		name      string
		mock      func()
		want      []models.ActorFilms
		wantError bool
	}{
		{
			name: "OK",
			mock: func() {
				rows := sqlmock.NewRows([]string{"actor_id", "name", "gender", "date_of_birth"}).
					AddRow(1, "John", "male", "1990-01-01").
					AddRow(2, "Emma", "female", "1985-05-12")
				mock.ExpectQuery("SELECT actor_id, name, gender, date_of_birth FROM actors").WillReturnRows(rows)
				mock.ExpectQuery("SELECT movie_id from movie_actors WHERE actor_id=?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"movie_id"}).AddRow(101))
				mock.ExpectQuery("SELECT movie_id from movie_actors WHERE actor_id=?").WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"movie_id"}).AddRow(102))
			},
			want: []models.ActorFilms{
				{
					Actor: models.Actor{
						ID:       1,
						Name:     "John",
						Gender:   "male",
						BirthDay: "1990-01-01",
					},
					Movies: []int{
						101,
					},
				},
				{
					Actor: models.Actor{
						ID:       2,
						Name:     "Emma",
						Gender:   "female",
						BirthDay: "1985-05-12",
					},
					Movies: []int{
						102,
					},
				},
			},
			wantError: false,
		},
		{
			name: "Database Query Error",
			mock: func() {
				mock.ExpectQuery("SELECT actor_id, name, gender, date_of_birth FROM actors").WillReturnError(errors.New("database error"))
			},
			wantError: true,
		},
		{
			name: "NoActors",
			mock: func() {
				rows := sqlmock.NewRows([]string{"actor_id", "name", "gender", "date_of_birth"})
				mock.ExpectQuery("SELECT actor_id, name, gender, date_of_birth FROM actors").WillReturnRows(rows)
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "NoMoviesForActor",
			mock: func() {
				rows := sqlmock.NewRows([]string{"actor_id", "name", "gender", "date_of_birth"}).
					AddRow(1, "John", "male", "1990-01-01").
					AddRow(2, "Emma", "female", "1985-05-12")
				mock.ExpectQuery("SELECT actor_id, name, gender, date_of_birth FROM actors").WillReturnRows(rows)
				mock.ExpectQuery("SELECT movie_id from movie_actors WHERE actor_id=?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"movie_id"})) // No movies for actor 1
				mock.ExpectQuery("SELECT movie_id from movie_actors WHERE actor_id=?").WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"movie_id"})) // No movies for actor 2
			},
			want: []models.ActorFilms{
				{
					Actor: models.Actor{
						ID:       1,
						Name:     "John",
						Gender:   "male",
						BirthDay: "1990-01-01",
					},
					Movies: nil,
				},
				{
					Actor: models.Actor{
						ID:       2,
						Name:     "Emma",
						Gender:   "female",
						BirthDay: "1985-05-12",
					},
					Movies: nil,
				},
			},
			wantError: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()
			got, err := r.GetActors()
			if testCase.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got)
			}
		})
	}
}
