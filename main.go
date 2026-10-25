package main

import (
	"log"
	"pos_material/config"
	"pos_material/database"
	"pos_material/handler"
	"pos_material/repositories"
	"pos_material/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	database.InitDB(cfg)

	// Inisialisasi Repository
	userRepo := repositories.NewUserRepository(database.DB)
	katRepo := repositories.NewKategoriRepository(database.DB)
	prodRepo := repositories.NewProdukRepository(database.DB)
	trxRepo := repositories.NewTransaksiRepository(database.DB)

	// Inisialisasi Service
	userService := services.NewUserService(userRepo)
	katService := services.NewKategoriService(katRepo)
	prodService := services.NewProdukService(prodRepo)
	trxService := services.NewTransaksiService(trxRepo)

	// Inisialisasi Handler
	userHandler := handler.NewUserHandler(userService)
	katHandler := handler.NewKategoriHandler(katService)
	prodHandler := handler.NewProdukHandler(prodService)
	trxHandler := handler.NewTransaksiHandler(trxService)

	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.Static("/static", "./static")

	// Web Page Routes
	r.GET("/", trxHandler.RenderLoginPage)
	r.GET("/transaksi", trxHandler.RenderTransaksiPage)
	r.GET("/produk", prodHandler.RenderProdukPage)
	r.GET("/laporan-produk", prodHandler.RenderReportProduk)
	r.GET("/laporan-transaksi", trxHandler.RenderReportTransaksi)
	r.GET("/user", userHandler.RenderUserPage)

	// API REST Routes
	api := r.Group("/api")
	{
		api.GET("/produk", prodHandler.ApiGetProduk)
		api.POST("/produk", prodHandler.ApiCreateProduk)
		api.POST("/kategori", katHandler.ApiCreateKategori)
		api.POST("/checkout", trxHandler.ApiCheckout)
		api.GET("/transaksi", trxHandler.ApiGetTransaksi)
		api.GET("/detail_penjualan", trxHandler.ApiGetDetailPenjualan)
		api.GET("/users", userHandler.ApiGetUsers)
	}

	log.Printf("Server berjalan di port %s", cfg.AppPort)
	r.Run(cfg.AppPort)
}