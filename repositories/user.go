package repositories

import (
	"pos_material/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAll(users *[]models.User) error {
	return r.db.Find(users).Error
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}