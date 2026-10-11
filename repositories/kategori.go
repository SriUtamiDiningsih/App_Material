package repositories

import (
	"pos_material/models"

	"gorm.io/gorm"
)

type KategoriRepository interface {
	FindAll() ([]models.Kategori, error)
	FindByID(id int) (*models.Kategori, error)
	Create(kategori *models.Kategori) error
	Update(kategori *models.Kategori) error
	Delete(id int) error
}

type kategoriRepository struct {
	db *gorm.DB
}

func NewKategoriRepository(db *gorm.DB) KategoriRepository {
	return &kategoriRepository{db: db}
}

func (r *kategoriRepository) FindAll() ([]models.Kategori, error) {
	var list []models.Kategori
	err := r.db.Order("kategori_id ASC").Find(&list).Error
	return list, err
}

func (r *kategoriRepository) FindByID(id int) (*models.Kategori, error) {
	var kat models.Kategori
	err := r.db.First(&kat, "kategori_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &kat, nil
}

func (r *kategoriRepository) Create(kategori *models.Kategori) error {
	return r.db.Create(kategori).Error
}

func (r *kategoriRepository) Update(kategori *models.Kategori) error {
	return r.db.Save(kategori).Error
}

func (r *kategoriRepository) Delete(id int) error {
	return r.db.Delete(&models.Kategori{}, "kategori_id = ?", id).Error
}