package actors_storage

import "github.com/moxicom/vk-internship-2024-spring/internal/models"

func (s *actorsStorage) AddActor(actor models.Actor) (int, error) {
	var id int
	err := s.db.QueryRow(
		"INSERT INTO actors(name, isMale, date_of_birth) VALUES($1, $2, $3) RETURNING actor_id",
		actor.Name,
		actor.IsMale,
		actor.BirthDay,
	).Scan(&id)
	if err != nil {
		return 0, nil
	}

	return id, nil
}
