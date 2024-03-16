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
