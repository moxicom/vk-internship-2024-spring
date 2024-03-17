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
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Discription string `json:"discription"` // Not required
	Date        string `json:"date" validate:"required"`
	Rating      *int   `json:"rating" validate:"required"` // 1-10
}

type SortParams struct {
	Sort  string
	Order string // asc or desc
}
