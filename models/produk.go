package models

type Produk struct {
	ProdukID   int      `gorm:"primaryKey;column:produk_id" json:"produk_id"`
	KategoriID *int     `gorm:"column:kategori_id" json:"kategori_id"`
	Kategori   Kategori `gorm:"foreignKey:KategoriID" json:"kategori"`
	Barcode    *string  `gorm:"size:50;unique" json:"barcode"`
	NamaProduk string   `gorm:"size:150;not null" json:"nama_produk"`
	Satuan     string   `gorm:"size:20;not null" json:"satuan"`
	HargaJual  float64  `gorm:"type:numeric(12,2);not null" json:"harga_jual"`
	Stok       int      `gorm:"default:0" json:"stok"`
	GambarURL  *string  `gorm:"size:255" json:"gambar_url"`
}

func (Produk) TableName() string {
	return "produks"
}