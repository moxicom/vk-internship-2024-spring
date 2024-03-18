package movies_storage_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres/movies_storage"
	"github.com/stretchr/testify/assert"
)

func TestMoviesStorage_GetMovie(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	r := movies_storage.New(db)

	type args struct {
		movieID int
	}

	type mockBehavior func(args args)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		result       models.MovieActors
		wantError    bool
	}{
		{
			name: "OK",
			args: args{
				movieID: 1,
			},
			mockBehavior: func(args args) {
				rows := sqlmock.NewRows([]string{"movie_id", "name", "description", "date", "rating"}).
					AddRow(args.movieID, "title", "", "2002-12-01", 5.0)
				mock.ExpectQuery("SELECT (.+) FROM movies").
					WithArgs(args.movieID).
					WillReturnRows(rows)
				mock.ExpectQuery("SELECT actor_id from movie_actors WHERE movie_id=?").
					WithArgs(args.movieID).
					WillReturnRows(
						sqlmock.NewRows([]string{"actor_id"}).AddRow(101),
					)
			},
			result: models.MovieActors{
				Movie: models.Movie{
					ID:          1,
					Name:        "title",
					Description: "",
					Date:        "2002-12-01",
					Rating:      &[]float32{5.0}[0],
				},
				Actors: []int{101},
			},
		},
		{
			name: "MovieNotFound",
			args: args{
				movieID: 1,
			},
			mockBehavior: func(args args) {
				mock.ExpectQuery("SELECT (.+) FROM movies").
					WithArgs(args.movieID).
					WillReturnError(errors.New("not found"))
			},
			wantError: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)

			res, err := r.GetMovie(testCase.args.movieID)
			if testCase.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.result, res)
			}

		})
	}
}
