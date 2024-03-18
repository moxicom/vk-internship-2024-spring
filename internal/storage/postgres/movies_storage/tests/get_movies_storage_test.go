package movies_storage_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres/movies_storage"
	"github.com/stretchr/testify/assert"
)

func TestMoviesStorage_GetMovies(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	r := movies_storage.New(db)

	type args struct {
		sort   models.SortParams
		search models.SearchParams
	}

	type mockBehavior func(args args)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		result       []models.MovieActors
		wantError    bool
	}{
		{
			name: "ok",
			args: args{},
			mockBehavior: func(args args) {
				mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{
					"movie_id",
					"name",
					"description",
					"date",
					"rating",
				}).
					AddRow(1, "interstellar", "", "2014-01-01", 9)
				mock.ExpectPrepare("SELECT").
					ExpectQuery().WithArgs().WillReturnRows(rows)
				mock.ExpectQuery("SELECT").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"actor_id"}).AddRow(2))

			},
			result: []models.MovieActors{
				models.MovieActors{
					Movie: models.Movie{
						ID:          1,
						Name:        "interstellar",
						Description: "",
						Date:        "2014-01-01",
						Rating:      &[]float32{9}[0],
					},
					Actors: []int{
						2,
					},
				},
			},
			wantError: false,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)

			res, err := r.GetMovies(testCase.args.sort, testCase.args.search)
			if testCase.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.result, res)
			}

		})
	}
}
