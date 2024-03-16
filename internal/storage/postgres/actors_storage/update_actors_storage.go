package actors_storage

import "github.com/moxicom/vk-internship-2024-spring/internal/models"

func (s *actorsStorage) UpdateActor(actorId int, actor models.ActorFilm) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		"UPDATE actors SET name = $1, isMale = $2, date_of_birth = $3 WHERE actor_id = $4",
		actor.Name,
		actor.Gender,
		actor.BirthDay,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
