package relations_storage_test

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres/relations_storage"
	"github.com/stretchr/testify/assert"
)

func TestRelationsStorage_DeleteRelation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := relations_storage.New(db)

	type mockBehavior func(args models.RelationMoviesActors)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		input        models.RelationMoviesActors
		wantError    bool
	}{
		{
			name: "OK",
			input: models.RelationMoviesActors{
				MovieID: "1",
				ActorID: "1",
			},
			mockBehavior: func(args models.RelationMoviesActors) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM movie_actors").
					WithArgs(args.MovieID, args.ActorID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantError: false,
		},
		{
			name: "Begin error",
			input: models.RelationMoviesActors{
				MovieID: "1",
				ActorID: "1",
			},
			mockBehavior: func(args models.RelationMoviesActors) {
				mock.ExpectBegin().WillReturnError(sqlmock.ErrCancelled)
			},
			wantError: true,
		},
		{
			name: "Exec error",
			input: models.RelationMoviesActors{
				MovieID: "1",
				ActorID: "1",
			},
			mockBehavior: func(args models.RelationMoviesActors) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM movie_actors").
					WithArgs(args.MovieID, args.ActorID).
					WillReturnError(sqlmock.ErrCancelled)
				mock.ExpectRollback()
			},
			wantError: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.input)
			err := r.DeleteRelation(testCase.input)
			if testCase.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
