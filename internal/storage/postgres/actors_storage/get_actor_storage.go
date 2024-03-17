package actors_storage

import "github.com/moxicom/vk-internship-2024-spring/internal/models"

func (s *actorsStorage) GetActor(actorId int) (models.ActorFilms, error) {
	actorRows, err := s.db.Query("SELECT actor_id, name, gender, date_of_birth FROM actors"+
		" WHERE actor_id = $1", actorId)
	if err != nil {
		return models.ActorFilms{}, err
	}
	defer actorRows.Close()

	var actor models.ActorFilms
	for actorRows.Next() {
		err := actorRows.Scan(&actor.ID, &actor.Name, &actor.Gender, &actor.BirthDay)
		if err != nil {
			return models.ActorFilms{}, err
		}
		movies, err := s.getActorFilmsIDs(actorId)
		if err != nil {
			return models.ActorFilms{}, nil
		}
		actor.Movies = movies
	}
	if err := actorRows.Err(); err != nil {
		return models.ActorFilms{}, err
	}
	return actor, nil
}
