package repositories

import (
	"fmt"

	"pos_material/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransaksiRepository interface {
	CreateTransactionACID(trx *models.Transaksi, details []models.DetailTransaksi) error
	FindAll() ([]models.Transaksi, error)
	FindByID(id int) (*models.Transaksi, error)
	FindByNoFaktur(noFaktur string) (*models.Transaksi, error)
}

type transaksiRepository struct {
	db *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) TransaksiRepository {
	return &transaksiRepository{db: db}
}

func (r *transaksiRepository) CreateTransactionACID(trx *models.Transaksi, details []models.DetailTransaksi) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for i := range details {
			var prod models.Produk
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&prod, "produk_id = ?", details[i].ProdukID).Error; err != nil {
				return fmt.Errorf("produk ID %d tidak ditemukan: %w", details[i].ProdukID, err)
			}

			if prod.Stok < details[i].Qty {
				return fmt.Errorf("stok produk '%s' tidak mencukupi (sisa: %d, diminta: %d)", prod.NamaProduk, prod.Stok, details[i].Qty)
			}

			newStock := prod.Stok - details[i].Qty
			if err := tx.Model(&models.Produk{}).Where("produk_id = ?", prod.ProdukID).Update("stok", newStock).Error; err != nil {
				return fmt.Errorf("gagal memotong stok produk %s: %w", prod.NamaProduk, err)
			}
		}

		if err := tx.Create(trx).Error; err != nil {
			return fmt.Errorf("gagal membuat nota transaksi: %w", err)
		}

		for i := range details {
			details[i].TransaksiID = trx.TransaksiID
			if err := tx.Create(&details[i]).Error; err != nil {
				return fmt.Errorf("gagal mencatat item transaksi ke-%d: %w", i+1, err)
			}
		}

		return nil
	})
}

func (r *transaksiRepository) FindAll() ([]models.Transaksi, error) {
	var list []models.Transaksi
	err := r.db.
		Preload("User").
		Preload("Pelanggan").
		Preload("DetailTransaksi.Produk").
		Order("transaksi_id DESC").
		Find(&list).Error
	return list, err
}

func (r *transaksiRepository) FindByID(id int) (*models.Transaksi, error) {
	var trx models.Transaksi
	err := r.db.
		Preload("User").
		Preload("Pelanggan").
		Preload("DetailTransaksi.Produk").
		First(&trx, "transaksi_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &trx, nil
}

func (r *transaksiRepository) FindByNoFaktur(noFaktur string) (*models.Transaksi, error) {
	var trx models.Transaksi
	err := r.db.
		Preload("User").
		Preload("Pelanggan").
		Preload("DetailTransaksi.Produk").
		First(&trx, "nomor_faktur = ?", noFaktur).Error
	if err != nil {
		return nil, err
	}
	return &trx, nil
}