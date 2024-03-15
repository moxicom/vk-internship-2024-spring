package handlers

import (
	"errors"
	"strconv"
)

func getIdByPrefix(idPath string) (int, error) {
	id, err := strconv.Atoi(idPath)
	if err != nil {
		return 0, errors.New("invalid id")
	}
	return id, nil
}
