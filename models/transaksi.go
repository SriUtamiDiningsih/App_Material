package models

import "time"

type Transaksi struct {
	TransaksiID      int               `gorm:"primaryKey;column:transaksi_id" json:"transaksi_id"`
	NomorFaktur      string            `gorm:"size:50;not null;unique" json:"nomor_faktur"`
	UserID           *int              `gorm:"column:user_id" json:"user_id"`
	User             User              `gorm:"foreignKey:UserID" json:"user"`
	PelangganID      *int              `gorm:"column:pelanggan_id" json:"pelanggan_id"`
	Pelanggan        Pelanggan         `gorm:"foreignKey:PelangganID" json:"pelanggan"`
	Subtotal         float64           `gorm:"type:numeric(12,2);not null" json:"subtotal"`
	PajakPersen      float64           `gorm:"type:numeric(5,2);default:11.00" json:"pajak_persen"`
	PajakNominal     float64           `gorm:"type:numeric(12,2);not null" json:"pajak_nominal"`
	TotalBayar       float64           `gorm:"type:numeric(12,2);not null" json:"total_bayar"`
	MetodePembayaran string            `gorm:"size:20;not null" json:"metode_pembayaran"`
	StatusTransaksi  string            `gorm:"size:20;default:Selesai" json:"status_transaksi"`
	CreatedAt        time.Time         `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	Details          []DetailTransaksi `gorm:"foreignKey:TransaksiID" json:"details"`
}

func (Transaksi) TableName() string {
	return "transaksi"
}