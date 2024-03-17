package actors_storage_test

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres/actors_storage"
	"github.com/stretchr/testify/assert"
)

func TestActorsStorage_DeleteActor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := actors_storage.New(db)

	type args struct {
		actorID int
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
				actorID: 1,
			},
			mockBehavior: func(args args) {
				mock.ExpectBegin()
				mock.ExpectExec("DELETE FROM movie_actors").WithArgs(args.actorID).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("DELETE FROM actors").WithArgs(args.actorID).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantError: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args)
			err := r.DeleteActor(testCase.args.actorID)
			if testCase.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
