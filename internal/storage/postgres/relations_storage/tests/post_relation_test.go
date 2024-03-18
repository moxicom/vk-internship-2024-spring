package relations_storage_test

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres/relations_storage"
)

func TestRelationsStorage_AddRelation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := relations_storage.New(db)

	type args struct {
		rel models.RelationMoviesActors
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
				rel: models.RelationMoviesActors{
					MovieID: "1",
					ActorID: "1",
				},
			},
			mockBehavior: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO movie_actors").
					WithArgs(args.rel.MovieID, args.rel.ActorID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantError: false,
		},
		{
			name: "Foreign key error",
			args: args{
				rel: models.RelationMoviesActors{
					MovieID: "1",
					ActorID: "1",
				},
			},
			mockBehavior: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT INTO movie_actors").
					WithArgs(args.rel.MovieID, args.rel.ActorID).
					WillReturnError(sqlmock.ErrCancelled)
				mock.ExpectRollback()
			},
			wantError: true,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior(tc.args)
			err := r.AddRelation(tc.args.rel)
			if (err != nil) != tc.wantError {
				t.Errorf("AddRelation() error = %v, wantError %v", err, tc.wantError)
				return
			}
		})
	}

}
