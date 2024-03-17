package actors_storage

import (
	"errors"
	"strconv"

	"github.com/moxicom/vk-internship-2024-spring/internal/models"
)

func (s *actorsStorage) UpdateActor(actorId int, actor models.Actor) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	queryParams := []interface{}{}
	updateQuery := "UPDATE actors SET"
	sqlArgIterator := 0

	if actor.Name != "" {
		sqlArgIterator++
		updateQuery += " name = $" + strconv.Itoa(sqlArgIterator) + ","
		queryParams = append(queryParams, actor.Name)
	}
	if actor.BirthDay != "" {
		sqlArgIterator++
		updateQuery += " date_of_birth = $" + strconv.Itoa(sqlArgIterator) + ","
		queryParams = append(queryParams, actor.BirthDay)
	}
	if actor.Gender != "" {
		sqlArgIterator++
		updateQuery += " gender = $" + strconv.Itoa(sqlArgIterator) + ","
		queryParams = append(queryParams, actor.Gender)
	}

	// Remove the trailing comma if any fields are updated
	if len(queryParams) > 0 {
		updateQuery = updateQuery[:len(updateQuery)-1]
	} else {
		return errors.New("no fields to update")
	}

	sqlArgIterator++
	updateQuery += " WHERE actor_id = $" + strconv.Itoa(sqlArgIterator)
	queryParams = append(queryParams, actorId)

	stmt, err := tx.Prepare(updateQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Execute the update actors
	_, err = stmt.Exec(queryParams...)
	if err != nil {
		return err
	}
	return tx.Commit()
}
