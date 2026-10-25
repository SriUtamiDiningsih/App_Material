package models

type Kategori struct {
	KategoriID   int    `gorm:"primaryKey;column:kategori_id" json:"kategori_id"`
	NamaKategori string `gorm:"size:50;not null" json:"nama_kategori"`
}

func (Kategori) TableName() string {
	return "kategori"
}