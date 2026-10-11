package models

type Produk struct {
	ProdukID   int      `gorm:"column:produk_id;primaryKey;autoIncrement" json:"produk_id"`
	KategoriID *int     `gorm:"column:kategori_id" json:"kategori_id"` // Gunakan *int karena nullable di DB
	Barcode    string   `gorm:"column:barcode;size:50" json:"barcode"`
	NamaProduk string   `gorm:"column:nama_produk;size:150;not null" json:"nama_produk"`
	Satuan     string   `gorm:"column:satuan;size:20;not null" json:"satuan"`
	HargaJual  float64  `gorm:"column:harga_jual;type:numeric(12,2);not null" json:"harga_jual"`
	Stok       int      `gorm:"column:stok;default:0" json:"stok"`
	GambarURL  string   `gorm:"column:gambar_url;size:255" json:"gambar_url"`

	// Relasi: Mengambil data kategori terkait (Belongs To)
	Kategori *Kategori `gorm:"foreignKey:KategoriID;references:KategoriID" json:"kategori,omitempty"`
}

// TableName mengunci nama tabel ke "produk"
func (Produk) TableName() string {
	return "produk"
}