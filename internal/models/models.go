package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type Actor struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	IsMale   bool   `json:"is_male"`
	BirthDay string `json:"birthday"`
}
