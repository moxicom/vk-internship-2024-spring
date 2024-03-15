package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type Actor struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	IsMale   bool   `json:"is_male"`
	BirthDay string `json:"birthday" validate:"required"`
}
