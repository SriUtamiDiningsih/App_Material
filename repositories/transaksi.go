package repositories

import (
	"pos_material/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransaksiRepository struct {
	db *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) *TransaksiRepository {
	return &TransaksiRepository{db: db}
}

func (r *TransaksiRepository) FindAllTransaksis(transaksis *[]models.Transaksi) error {
	return r.db.Preload("User").Preload("Pelanggan").Find(transaksis).Error
}

func (r *TransaksiRepository) FindAllDetails(details *[]models.DetailTransaksi) error {
	return r.db.Preload("Produk").Find(details).Error
}

// Atomic Transaction untuk Kasir (Scan, kurangi stok, simpan detail, cetak faktur)
func (r *TransaksiRepository) CreateCheckout(trx *models.Transaksi) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(trx).Error; err != nil {
			return err
		}

		for _, detail := range trx.Details {
			detail.TransaksiID = trx.TransaksiID
			if err := tx.Create(&detail).Error; err != nil {
				return err
			}

			// Validasi & Kurangi Stok Produk secara atomik
			var produk models.Produk
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("produk_id = ?", detail.ProdukID).First(&produk).Error; err != nil {
				return err
			}

			if produk.Stok < detail.Qty {
				return gorm.ErrInvalidData
			}

			produk.Stok -= detail.Qty
			if err := tx.Save(&produk).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
