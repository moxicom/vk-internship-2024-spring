package users_storage

import "database/sql"

type UsersStorage struct {
	db *sql.DB
}

func New(db *sql.DB) *UsersStorage {
	return &UsersStorage{db}
}

func (s *UsersStorage) CheckUser(username, password string, isAdmin bool) (bool, error) {
	var count int
	err := s.db.QueryRow(
		"SELECT COUNT(*) FROM users WHERE username=$1"+
			" AND password_hash=$2 AND isAdmin=$3",
		username,
		password,
		isAdmin).Scan(&count)

	if err != nil {
		return false, err
	}
	return count > 0, nil
}
