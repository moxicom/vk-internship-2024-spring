package movies_storage_test

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres/movies_storage"
	"github.com/stretchr/testify/assert"
)

func TestMoviesStorage_DeleteMovie(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
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
		wantError    bool
	}{
		{
			name: "OK",
			args: args{
				movieID: 1,
			},
			mockBehavior: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM movie_actors").
					WithArgs(args.movieID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("DELETE FROM movies").
					WithArgs(args.movieID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantError: false,
		},
		{
			name: "Begin error",
			args: args{
				movieID: 1,
			},
			mockBehavior: func(args args) {
				mock.ExpectBegin().WillReturnError(sqlmock.ErrCancelled)
			},
			wantError: true,
		},
		{
			name: "Delete movie error",
			args: args{
				movieID: 1,
			},
			mockBehavior: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM movie_actors").
					WithArgs(args.movieID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("DELETE FROM movies").
					WithArgs(args.movieID).
					WillReturnError(sqlmock.ErrCancelled)
			},
			wantError: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)
			err := r.DeleteMovie(testCase.args.movieID)
			if testCase.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
