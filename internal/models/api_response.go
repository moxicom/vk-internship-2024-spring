package models

// Actor and its films id
type ActorFilm struct {
	Actor
	Movies []int `json:"movies"`
}
