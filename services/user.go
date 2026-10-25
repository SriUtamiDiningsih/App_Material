package services

import (
	"pos_material/models"
	"pos_material/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers(users *[]models.User) error {
	return s.repo.FindAll(users)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}