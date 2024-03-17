package actors_storage

import "github.com/moxicom/vk-internship-2024-spring/internal/models"

func (s *actorsStorage) AddActor(actor models.Actor) (int, error) {
	var id int
	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}

	err = tx.QueryRow(
		"INSERT INTO actors(name, gender, date_of_birth) VALUES($1, $2, $3) RETURNING actor_id",
		actor.Name,
		actor.Gender,
		actor.BirthDay,
	).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, nil
	}

	return id, tx.Commit()
}
