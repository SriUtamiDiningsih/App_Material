package repositories

import (
	"pos_material/models"
	"gorm.io/gorm"
)

type ProdukRepository struct {
	db *gorm.DB
}

func NewProdukRepository(db *gorm.DB) *ProdukRepository {
	return &ProdukRepository{db: db}
}

func (r *ProdukRepository) FindAll(produks *[]models.Produk) error {
	return r.db.Preload("Kategori").Find(produks).Error
}

func (r *ProdukRepository) Create(produk *models.Produk) error {
	return r.db.Create(produk).Error
}