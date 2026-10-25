package repositories

import (
	"pos_material/models"
	"gorm.io/gorm"
)

type KategoriRepository struct {
	db *gorm.DB
}

func NewKategoriRepository(db *gorm.DB) *KategoriRepository {
	return &KategoriRepository{db: db}
}

func (r *KategoriRepository) FindAll(kategoris *[]models.Kategori) error {
	return r.db.Find(kategoris).Error
}

func (r *KategoriRepository) Create(kategori *models.Kategori) error {
	return r.db.Create(kategori).Error
}