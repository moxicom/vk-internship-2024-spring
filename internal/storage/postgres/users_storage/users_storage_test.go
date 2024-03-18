package users_storage_test

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres/users_storage"
)

func TestUsersStorage_CkeckUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := users_storage.New(db)

	type args struct {
		username string
		password string
		isAdmin  bool
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
				username: "user",
				password: "password",
				isAdmin:  true,
			},
			mockBehavior: func(args args) {
				mock.ExpectQuery("SELECT COUNT").
					WithArgs(args.username, args.password, args.isAdmin).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).
						AddRow(1),
					)
			},
			wantError: false,
		},
		{
			name: "Error",
			args: args{
				username: "user",
				password: "password",
				isAdmin:  true,
			},
			mockBehavior: func(args args) {
				mock.ExpectQuery("SELECT COUNT").
					WithArgs(args.username, args.password, args.isAdmin).
					WillReturnError(sqlmock.ErrCancelled)
			},
			wantError: true,
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockBehavior(tc.args)
			_, err := r.CheckUser(tc.args.username, tc.args.password, tc.args.isAdmin)
			if (err != nil) != tc.wantError {
				t.Errorf("error = %v, wantError %v", err, tc.wantError)
				return
			}
		})
	}
}
