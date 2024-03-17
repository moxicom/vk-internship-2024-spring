package utils

import (
	"errors"

	"github.com/moxicom/vk-internship-2024-spring/internal/models"
)

func ValidateSortParams(sort models.SortParams) error {
	// CHECK IF THERE IS NO SORT PARAMS
	if sort.Sort == "" && sort.Order == "" {
		return nil
	}
	err := errors.New("invalid sort params")
	flag := false
	switch sort.Sort {
	case "name":
		flag = true
	case "date":
		flag = true
	case "rating":
		flag = true
	}
	if !flag {
		return err
	}
	if sort.Order != "asc" && sort.Order != "desc" && sort.Order != "" {
		return err
	}
	return nil
}
