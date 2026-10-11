package models

import "time"

type User struct {
	UserID      int       `gorm:"column:user_id;primaryKey;autoIncrement" json:"user_id"`
	Username    string    `gorm:"column:username;size:50;not null" json:"username"`
	NamaLengkap string    `gorm:"column:nama_lengkap;size:100;not null" json:"nama_lengkap"`
	Role        string    `gorm:"column:role;size:20;default:'kasir'" json:"role"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName mengunci nama tabel ke "users"
func (User) TableName() string {
	return "users"
}