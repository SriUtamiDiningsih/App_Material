package models

import "time"

type Pelanggan struct {
	PelangganID   int       `gorm:"primaryKey;column:pelanggan_id" json:"pelanggan_id"`
	KodeMember    string    `gorm:"size:20;not null;unique" json:"kode_member"`
	NamaPelanggan string    `gorm:"size:100;not null" json:"nama_pelanggan"`
	NoTelepon     *string   `gorm:"size:15" json:"no_telepon"`
	Alamat        *string   `gorm:"type:text" json:"alamat"`
	CreatedAt     time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (Pelanggan) TableName() string {
	return "pelanggan"
}