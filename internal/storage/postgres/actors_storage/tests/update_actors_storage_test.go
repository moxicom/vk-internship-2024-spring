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

func TestActorsStorage_UpdateActor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := actors_storage.New(db)

	type args struct {
		actorId int
		actor   models.Actor
	}

	type mockBehavior func(args args, id int)

	testTable := []struct {
		name         string
		args         args
		mockBehavior mockBehavior
		wantError    bool
	}{
		{
			name: "OK",
			args: args{
				actorId: 1,
				actor: models.Actor{
					Name:   "Alexander",
					Gender: "male",
				},
			},
			mockBehavior: func(args args, id int) {
				mock.ExpectBegin()
				mock.ExpectPrepare("UPDATE").ExpectExec().WithArgs(args.actor.Name, args.actor.Gender, args.actorId).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantError: false,
		},
		{
			name: "No fields to update",
			args: args{
				actorId: 1,
				actor:   models.Actor{},
			},
			mockBehavior: func(args args, id int) {
				mock.ExpectBegin()
				mock.ExpectPrepare("UPDATE").ExpectExec().WithArgs(args.actor.Name, args.actor.Gender, args.actorId).WillReturnError(errors.New("no fields to update"))
				mock.ExpectRollback()
			},
			wantError: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args, 1)
			err := r.UpdateActor(testCase.args.actorId, testCase.args.actor)
			if testCase.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
