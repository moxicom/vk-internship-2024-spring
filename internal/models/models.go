package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type Actor struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Gender   string `json:"gender" validate:"required"`
	BirthDay string `json:"birthday" validate:"required"`
}

type Movie struct {
	ID          int      `json:"id"`
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description"` // Not required
	Date        string   `json:"date" validate:"required"`
	Rating      *float32 `json:"rating" validate:"required"` // 1-10
}

type SortParams struct {
	Sort  string
	Order string // asc or desc
}

type SearchParams struct {
	MovieName string
	ActorName string
}

type RelationMoviesActors struct {
	MovieID string `json:"movie_id" validate:"required"`
	ActorID string `json:"actor_id" validate:"required"`
}
