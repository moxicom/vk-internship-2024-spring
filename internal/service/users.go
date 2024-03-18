package service

import (
	"github.com/moxicom/vk-internship-2024-spring/internal/storage"
)

type usersService struct {
	storage *storage.Repository
}

func newUsersService(s *storage.Repository) *usersService {
	return &usersService{s}
}

func (s *usersService) CheckUser(username string, password string, isAdmin bool) (bool, error) {
	return s.storage.Users.CheckUser(username, password, isAdmin)
}
