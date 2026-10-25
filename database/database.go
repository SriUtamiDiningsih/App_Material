package database

import (
	"fmt"
	"log"
	"pos_material/config"
	"pos_material/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database PostgreSQL: %v", err)
	}

	log.Println("Berhasil terhubung ke database PostgreSQL!")

	// Auto Migrate Semua Model Domain
	err = DB.AutoMigrate(
		&models.User{},
		&models.Kategori{},
		&models.Pelanggan{},
		&models.Produk{},
		&models.Transaksi{},
		&models.DetailTransaksi{},
	)
	if err != nil {
		log.Fatalf("Gagal migrasi database: %v", err)
	}
}