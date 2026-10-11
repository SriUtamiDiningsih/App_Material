package models

import "time"

type Transaksi struct {
	TransaksiID      int       `gorm:"column:transaksi_id;primaryKey;autoIncrement" json:"transaksi_id"`
	NomorFaktur      string    `gorm:"column:nomor_faktur;size:50;not null" json:"nomor_faktur"`
	UserID           *int      `gorm:"column:user_id" json:"user_id"`           // nullable jika tidak login/kasir default
	PelangganID      *int      `gorm:"column:pelanggan_id" json:"pelanggan_id"` // nullable untuk pembeli non-member
	Subtotal         float64   `gorm:"column:subtotal;type:numeric(12,2);not null" json:"subtotal"`
	PajakPersen      float64   `gorm:"column:pajak_persen;type:numeric(5,2);default:11.00" json:"pajak_persen"`
	PajakNominal     float64   `gorm:"column:pajak_nominal;type:numeric(12,2);not null" json:"pajak_nominal"`
	TotalBayar       float64   `gorm:"column:total_bayar;type:numeric(12,2);not null" json:"total_bayar"`
	MetodePembayaran string    `gorm:"column:metode_pembayaran;size:20" json:"metode_pembayaran"`
	StatusTransaksi  string    `gorm:"column:status_transaksi;size:20;default:'Selesai'" json:"status_transaksi"`
	CreatedAt        time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`

	// Relasi
	User            *User             `gorm:"foreignKey:UserID;references:UserID" json:"user,omitempty"`
	Pelanggan       *Pelanggan        `gorm:"foreignKey:PelangganID;references:PelangganID" json:"pelanggan,omitempty"`
	DetailTransaksi []DetailTransaksi `gorm:"foreignKey:TransaksiID;references:TransaksiID" json:"detail_transaksi,omitempty"`
}

// TableName mengunci nama tabel ke "transaksi"
func (Transaksi) TableName() string {
	return "transaksi"
}