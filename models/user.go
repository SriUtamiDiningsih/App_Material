package models

import "time"

type User struct {
	UserID      int       `gorm:"primaryKey;column:user_id" json:"user_id"`
	Username    string    `gorm:"size:50;not null;unique" json:"username"`
	Password    string    `gorm:"size:255;not null" json:"password"`
	NamaLengkap string    `gorm:"size:100;not null" json:"nama_lengkap"`
	Role        string    `gorm:"size:20;default:kasir" json:"role"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (User) TableName() string {
	return "users"
}