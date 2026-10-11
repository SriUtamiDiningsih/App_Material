package repositories

import (
	"pos_material/models"

	"gorm.io/gorm"
)

type ProdukRepository interface {
	FindAll(search string, kategoriID *int) ([]models.Produk, error)
	FindByID(id int) (*models.Produk, error)
	FindByBarcode(barcode string) (*models.Produk, error)
	Create(produk *models.Produk) error
	Update(produk *models.Produk) error
	Delete(id int) error
	CountTotal() (int64, error)
	CountInStock() (int64, error)
	CountLowStock(threshold int) (int64, error)
}

type produkRepository struct {
	db *gorm.DB
}

func NewProdukRepository(db *gorm.DB) ProdukRepository {
	return &produkRepository{db: db}
}

func (r *produkRepository) FindAll(search string, kategoriID *int) ([]models.Produk, error) {
	var list []models.Produk
	query := r.db.Preload("Kategori").Order("produk_id ASC")

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("nama_produk ILIKE ? OR barcode ILIKE ?", s, s)
	}

	if kategoriID != nil && *kategoriID > 0 {
		query = query.Where("kategori_id = ?", *kategoriID)
	}

	err := query.Find(&list).Error
	return list, err
}

func (r *produkRepository) FindByID(id int) (*models.Produk, error) {
	var p models.Produk
	err := r.db.Preload("Kategori").First(&p, "produk_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *produkRepository) FindByBarcode(barcode string) (*models.Produk, error) {
	var p models.Produk
	err := r.db.Preload("Kategori").First(&p, "barcode = ?", barcode).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *produkRepository) Create(produk *models.Produk) error {
	return r.db.Create(produk).Error
}

func (r *produkRepository) Update(produk *models.Produk) error {
	return r.db.Save(produk).Error
}

func (r *produkRepository) Delete(id int) error {
	return r.db.Delete(&models.Produk{}, "produk_id = ?", id).Error
}

func (r *produkRepository) CountTotal() (int64, error) {
	var count int64
	err := r.db.Model(&models.Produk{}).Count(&count).Error
	return count, err
}

func (r *produkRepository) CountInStock() (int64, error) {
	var count int64
	err := r.db.Model(&models.Produk{}).Where("stok > 0").Count(&count).Error
	return count, err
}

func (r *produkRepository) CountLowStock(threshold int) (int64, error) {
	var count int64
	err := r.db.Model(&models.Produk{}).Where("stok <= ?", threshold).Count(&count).Error
	return count, err
}