package utils

import (
	"strings"

	"github.com/moxicom/vk-internship-2024-spring/internal/models"
)

func ProcessSearchParams(search models.SearchParams) models.SearchParams {
	if search.ActorName != "" {
		search.ActorName = strings.ToLower(search.ActorName)
	}
	if search.MovieName != "" {
		search.MovieName = strings.ToLower(search.MovieName)
	}
	return search
}
