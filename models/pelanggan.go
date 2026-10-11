package models

import "time"

type Pelanggan struct {
	PelangganID   int       `gorm:"column:pelanggan_id;primaryKey;autoIncrement" json:"pelanggan_id"`
	KodeMember    string    `gorm:"column:kode_member;size:20;not null" json:"kode_member"`
	NamaPelanggan string    `gorm:"column:nama_pelanggan;size:100;not null" json:"nama_pelanggan"`
	NoTelepon     string    `gorm:"column:no_telepon;size:15" json:"no_telepon"`
	Alamat        string    `gorm:"column:alamat;type:text" json:"alamat"`
	CreatedAt     time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName mengunci nama tabel ke "pelanggan"
func (Pelanggan) TableName() string {
	return "pelanggan"
}