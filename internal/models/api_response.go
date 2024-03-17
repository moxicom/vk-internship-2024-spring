package models

// Actor and its films id
type ActorFilms struct {
	Actor
	Movies []int `json:"movies"`
}

type MovieActors struct {
	Movie
	Actors []int `json:"actors"`
}
