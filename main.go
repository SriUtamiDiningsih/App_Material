package main

import (
	"fmt"
	"log"
	"net/http"

	"pos_material/config"
	"pos_material/database"
	"pos_material/handler"
	"pos_material/repositories"
	"pos_material/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Memuat konfigurasi dari file .env
	cfg := config.LoadConfig()

	// 2. Inisialisasi koneksi PostgreSQL ke db_material
	db := database.InitDB(cfg)

	// 3. Inisialisasi Lapisan Repository (Data Access Layer)
	kategoriRepo := repositories.NewKategoriRepository(db)
	produkRepo := repositories.NewProdukRepository(db)
	transaksiRepo := repositories.NewTransaksiRepository(db)
	userRepo := repositories.NewUserRepository(db) // <-- DITAMBAHKAN UNTUK USER

	// 4. Inisialisasi Lapisan Service (Business Logic Layer)
	kategoriService := services.NewKategoriService(kategoriRepo)
	produkService := services.NewProdukService(produkRepo)
	transaksiService := services.NewTransaksiService(transaksiRepo, produkRepo)
	userService := services.NewUserService(userRepo) // <-- DITAMBAHKAN UNTUK USER

	// 5. Inisialisasi Lapisan Handler (Controller API)
	kategoriHandler := handler.NewKategoriHandler(kategoriService)
	produkHandler := handler.NewProdukHandler(produkService)
	transaksiHandler := handler.NewTransaksiHandler(transaksiService)
	userHandler := handler.NewUserHandler(userService) // <-- DITAMBAHKAN UNTUK USER

	// 6. Setup Router Gin
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Melayani Berkas Statis (CSS, JS, Gambar)
	r.Static("/static", "./static")

	// -------------------------------------------------------------
	// Rute Halaman Web (HTML Template)
	// -------------------------------------------------------------
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/produk")
	})
	r.GET("/login", func(c *gin.Context) {
		c.File("template/login.html")
	})
	r.GET("/produk", func(c *gin.Context) {
		c.File("template/produk.html")
	})
	r.GET("/transaksi", func(c *gin.Context) {
		c.File("template/transaksi.html")
	})
	r.GET("/pelanggan", func(c *gin.Context) {
		c.File("template/pelanggan.html")
	})
	r.GET("/user", func(c *gin.Context) {
		c.File("template/user.html")
	})
	r.GET("/laporan/produk", func(c *gin.Context) {
		c.File("template/laporan_produk.html")
	})
	r.GET("/laporan/penjualan", func(c *gin.Context) {
		c.File("template/laporan_penjualan.html")
	})

	// -------------------------------------------------------------
	// Rute REST API
	// -------------------------------------------------------------
	api := r.Group("/api")
	{
		// Endpoints Produk
		api.GET("/produk", produkHandler.GetAll)
		api.GET("/produk/stats", produkHandler.GetStats)
		api.GET("/produk/:id", produkHandler.GetByID)
		api.GET("/produk/barcode/:barcode", produkHandler.GetByBarcode)
		api.POST("/produk", produkHandler.Create)
		api.PUT("/produk/:id", produkHandler.Update)
		api.DELETE("/produk/:id", produkHandler.Delete)

		// Endpoints Kategori
		api.GET("/kategori", kategoriHandler.GetAll)
		api.GET("/kategori/:id", kategoriHandler.GetByID)
		api.POST("/kategori", kategoriHandler.Create)
		api.PUT("/kategori/:id", kategoriHandler.Update)
		api.DELETE("/kategori/:id", kategoriHandler.Delete)

		// Endpoints Transaksi Kasir (ACID Transaction)
		api.POST("/transaksi/checkout", transaksiHandler.Checkout)
		api.GET("/transaksi", transaksiHandler.GetAll)
		api.GET("/transaksi/:id", transaksiHandler.GetByID)
		api.GET("/transaksi/faktur/:no_faktur", transaksiHandler.GetByNoFaktur)

		// Endpoints User Pengguna (DITAMBAHKAN)
		api.GET("/users", userHandler.GetAll)
		api.POST("/users", userHandler.Create)
		api.DELETE("/users/:id", userHandler.Delete)
	}

	// 7. Menjalankan Server
	serverAddr := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("🚀 Server Toko Material POS berjalan di http://localhost%s\n", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}