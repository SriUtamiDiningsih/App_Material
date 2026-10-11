package services

import (
	"errors"
	"strings"

	"pos_material/models"
	"pos_material/repositories"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	CreateUser(u *models.User) error
	DeleteUser(id int) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id int) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) CreateUser(u *models.User) error {
	u.Username = strings.TrimSpace(u.Username)
	u.NamaLengkap = strings.TrimSpace(u.NamaLengkap)

	if u.Username == "" {
		return errors.New("username wajib diisi")
	}
	if u.NamaLengkap == "" {
		return errors.New("nama lengkap wajib diisi")
	}
	if u.Role == "" {
		u.Role = "kasir"
	}

	existing, _ := s.repo.FindByUsername(u.Username)
	if existing != nil {
		return errors.New("username sudah digunakan, gunakan username lain")
	}

	return s.repo.Create(u)
}

func (s *userService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}