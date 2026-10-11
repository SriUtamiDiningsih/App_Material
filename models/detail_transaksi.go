package models

type DetailTransaksi struct {
	DetailID    int     `gorm:"column:detail_id;primaryKey;autoIncrement" json:"detail_id"`
	TransaksiID int     `gorm:"column:transaksi_id;not null" json:"transaksi_id"`
	ProdukID    int     `gorm:"column:produk_id;not null" json:"produk_id"`
	HargaSatuan float64 `gorm:"column:harga_satuan;type:numeric(12,2);not null" json:"harga_satuan"`
	Qty         int     `gorm:"column:qty;not null" json:"qty"`
	Subtotal    float64 `gorm:"column:subtotal;type:numeric(12,2);not null" json:"subtotal"`

	// Relasi Opsional (untuk Preload data produk saat query nota/invoice)
	Produk *Produk `gorm:"foreignKey:ProdukID;references:ProdukID" json:"produk,omitempty"`
}

// TableName mengunci nama tabel agar GORM tidak mencari "detail_transaksis"
func (DetailTransaksi) TableName() string {
	return "detail_transaksi"
}