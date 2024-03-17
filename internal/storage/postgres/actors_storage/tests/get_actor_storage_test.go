package actors_storage

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres/actors_storage"
	"github.com/stretchr/testify/assert"
)

func TestActorsStorage_GetActor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := actors_storage.New(db)

	type args struct {
		actorId int
	}

	tests := []struct {
		name      string
		args      args
		mock      func(actorId int)
		want      models.ActorFilm
		wantError bool
	}{
		{
			name: "OK",
			args: args{1},
			mock: func(actorId int) {
				actorRows := sqlmock.NewRows([]string{"actor_id", "name", "gender", "date_of_birth"}).
					AddRow(1, "John", "male", "1990-01-01")

				mock.ExpectQuery("SELECT actor_id, name, gender, date_of_birth FROM actors").
					WithArgs(actorId).
					WillReturnRows(actorRows)
				mock.ExpectQuery("SELECT movie_id from movie_actors WHERE actor_id=?").WithArgs(actorId).WillReturnRows(sqlmock.NewRows([]string{"movie_id"}).AddRow(101))
			},
			want: models.ActorFilm{
				Actor: models.Actor{
					ID:       1,
					Name:     "John",
					Gender:   "male",
					BirthDay: "1990-01-01",
				},
				Movies: []int{
					101,
				},
			},
			wantError: false,
		},
		{
			name: "ErrorActorNotFound",
			args: args{1},
			mock: func(actorId int) {
				mock.ExpectQuery("SELECT actor_id, name, gender, date_of_birth FROM actors").
					WithArgs(actorId).
					WillReturnError(sql.ErrNoRows)
			},
			wantError: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock(testCase.args.actorId)
			got, err := r.GetActor(testCase.args.actorId)
			if testCase.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got)
			}
		})
	}

}
