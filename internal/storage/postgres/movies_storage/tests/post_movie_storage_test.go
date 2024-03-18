package movies_storage_test

import (
	"errors"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres/movies_storage"
	"github.com/stretchr/testify/assert"
)

func TestMoviesStorage_AddMovie(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := movies_storage.New(db)

	type args struct {
		movie models.Movie
	}

	type mockBehavior func(args args)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		expected     int
		wantError    bool
	}{
		{
			name: "OK",
			args: args{
				movie: models.Movie{
					Name:        "Movie",
					Description: "Description",
					Date:        "2022-12-12",
					Rating:      &[]float32{5.0}[0],
				},
			},
			expected: 1,
			mockBehavior: func(args args) {
				mock.ExpectBegin()
				rows := mock.NewRows([]string{"movie_id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO movies").
					WithArgs(args.movie.Name, args.movie.Description, args.movie.Date, args.movie.Rating).
					WillReturnRows(rows)
				mock.ExpectCommit()
			},
			wantError: false,
		},
		{
			name: "Insert error",
			args: args{
				movie: models.Movie{
					Name:        "Movie",
					Description: "Description",
					Date:        "2022-12-12",
					Rating:      &[]float32{5.0}[0],
				},
			},
			expected: 1,
			mockBehavior: func(args args) {
				mock.ExpectBegin()
				mock.ExpectQuery("INSERT INTO movies").
					WithArgs(args.movie.Name, args.movie.Description, args.movie.Date, args.movie.Rating).
					WillReturnError(errors.New("insert error"))
			},
			wantError: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)
			got, err := r.AddMovie(testCase.args.movie)
			if testCase.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expected, got)
			}
		})
	}
}
