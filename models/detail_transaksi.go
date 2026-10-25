package models

type DetailTransaksi struct {
	DetailID    int     `gorm:"primaryKey;column:detail_id" json:"detail_id"`
	TransaksiID int     `gorm:"column:transaksi_id;not null" json:"transaksi_id"`
	ProdukID    int     `gorm:"column:produk_id;not null" json:"produk_id"`
	Produk      Produk  `gorm:"foreignKey:ProdukID" json:"produk"`
	HargaSatuan float64 `gorm:"type:numeric(12,2);not null" json:"harga_satuan"`
	Qty         int     `gorm:"not null" json:"qty"`
	Subtotal    float64 `gorm:"type:numeric(12,2);not null" json:"subtotal"`
}

func (DetailTransaksi) TableName() string {
	return "detail_penjualan"
}