package models

type Kategori struct {
	KategoriID   int      `gorm:"column:kategori_id;primaryKey;autoIncrement" json:"kategori_id"`
	NamaKategori string   `gorm:"column:nama_kategori;size:50;not null" json:"nama_kategori"`

	// Relasi: Satu kategori bisa memiliki banyak produk (Has Many)
	Produk []Produk `gorm:"foreignKey:KategoriID;references:KategoriID" json:"produk,omitempty"`
}

// TableName mengunci nama tabel ke "kategori"
func (Kategori) TableName() string {
	return "kategori"
}