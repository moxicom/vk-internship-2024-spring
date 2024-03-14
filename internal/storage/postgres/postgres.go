package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
)

type Storage struct {
	db *sql.DB
}

func New(postgresConfig models.PostgresConfig) (*Storage, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		postgresConfig.PostgresHost,
		postgresConfig.PostgresPort,
		postgresConfig.PostgresUser,
		postgresConfig.PostgresPassword,
		postgresConfig.PostgresName,
		postgresConfig.PostgresSSLMode,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return &Storage{}, err
	}
	if err = db.Ping(); err != nil {
		return &Storage{}, err
	}
	_, err = db.Query("SELECT * FROM users;")
	if err != nil {
		return &Storage{}, err
	}
	return &Storage{db}, nil
}
