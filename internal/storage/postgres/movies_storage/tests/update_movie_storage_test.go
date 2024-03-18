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

func TestMoviesStorage_UpdateMovie(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := movies_storage.New(db)

	type args struct {
		movieId int
		movie   models.Movie
	}

	type mockBehavior func(args args)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		wantError    bool
	}{
		{
			name: "OK",
			args: args{
				movie: models.Movie{
					Name: "New name",
				},
			},
			mockBehavior: func(args args) {
				mock.ExpectBegin()
				mock.ExpectPrepare("UPDATE").
					ExpectExec().
					WithArgs(args.movie.Name, args.movieId).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantError: false,
		},
		{
			name: "UpdateErr",
			args: args{
				movie: models.Movie{
					Name: "New name",
				},
			},
			mockBehavior: func(args args) {
				mock.ExpectBegin()
				mock.ExpectPrepare("UPDATE").
					ExpectExec().
					WithArgs(args.movie.Name, args.movieId).
					WillReturnError(errors.New("insert error"))
			},
			wantError: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)
			err := r.UpdateMovie(testCase.args.movieId, testCase.args.movie)
			if testCase.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
