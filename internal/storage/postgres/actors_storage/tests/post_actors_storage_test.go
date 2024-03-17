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

func TestActorsStorage_AddActor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := actors_storage.New(db)

	type args struct {
		actor models.Actor
	}

	type mockBehavior func(args args, id int)

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
				actor: models.Actor{
					Name:     "Alexander",
					Gender:   "male",
					BirthDay: "2004-14-12",
				},
			},
			expected: 1,
			mockBehavior: func(args args, id int) {
				mock.ExpectBegin()
				rows := mock.NewRows([]string{"actor_id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO actors").WithArgs(args.actor.Name, args.actor.Gender, args.actor.BirthDay).WillReturnRows(rows)
				mock.ExpectCommit()
			},
		},
		{
			name: "Begin Transaction Error",
			mockBehavior: func(args args, id int) {
				mock.ExpectBegin().WillReturnError(errors.New("begin transaction error"))
			},
			wantError: true,
		},
		{
			name: "Valid Female Actor",
			args: args{
				actor: models.Actor{
					Name:     "Emma",
					Gender:   "female", // Female actor
					BirthDay: "1990-05-20",
				},
			},
			expected: 2, // Assuming the actor_id for this entry is 3
			mockBehavior: func(args args, id int) {
				mock.ExpectBegin()
				rows := mock.NewRows([]string{"actor_id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO actors").WithArgs(args.actor.Name, args.actor.Gender, args.actor.BirthDay).WillReturnRows(rows)
				mock.ExpectCommit()
			},
		},
		{
			name: "Null Birth Date",
			args: args{
				actor: models.Actor{
					Name:     "James",
					Gender:   "male",
					BirthDay: "", // Null birth date
				},
			},
			expected: 3, // Assuming the actor_id for this entry is 4
			mockBehavior: func(args args, id int) {
				mock.ExpectBegin()
				rows := mock.NewRows([]string{"actor_id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO actors").WithArgs(args.actor.Name, args.actor.Gender, args.actor.BirthDay).WillReturnRows(rows)
				mock.ExpectCommit()
			},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args, testCase.expected)
			got, err := r.AddActor(testCase.args.actor)
			if testCase.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expected, got)
			}
		})
	}
}
