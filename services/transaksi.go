package services

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"

	"pos_material/models"
	"pos_material/repositories"
)

type CheckoutItem struct {
	ProdukID    int     `json:"produk_id" binding:"required"`
	Qty         int     `json:"qty" binding:"required,min=1"`
	HargaSatuan float64 `json:"harga_satuan" binding:"required,min=0"`
	Diskon      float64 `json:"diskon"`
}

type CheckoutRequest struct {
	UserID           *int           `json:"user_id"`
	PelangganID      *int           `json:"pelanggan_id"`
	MetodePembayaran string         `json:"metode_pembayaran"`
	Items            []CheckoutItem `json:"items" binding:"required,min=1"`
	UangBayar        float64        `json:"uang_bayar"`
}

type CheckoutResponse struct {
	Transaksi *models.Transaksi `json:"transaksi"`
	Kembalian float64           `json:"kembalian"`
}

type TransaksiService interface {
	ProcessCheckout(req CheckoutRequest) (*CheckoutResponse, error)
	GetAllTransaksi() ([]models.Transaksi, error)
	GetTransaksiByID(id int) (*models.Transaksi, error)
	GetTransaksiByNoFaktur(noFaktur string) (*models.Transaksi, error)
}

type transaksiService struct {
	transaksiRepo repositories.TransaksiRepository
	produkRepo    repositories.ProdukRepository
}

func NewTransaksiService(
	transaksiRepo repositories.TransaksiRepository,
	produkRepo repositories.ProdukRepository,
) TransaksiService {
	return &transaksiService{
		transaksiRepo: transaksiRepo,
		produkRepo:    produkRepo,
	}
}

func (s *transaksiService) ProcessCheckout(req CheckoutRequest) (*CheckoutResponse, error) {
	if len(req.Items) == 0 {
		return nil, errors.New("keranjang transaksi tidak boleh kosong")
	}

	var subtotal float64 = 0
	var detailList []models.DetailTransaksi

	for _, it := range req.Items {
		itemSubtotal := (it.HargaSatuan * float64(it.Qty)) - it.Diskon
		if itemSubtotal < 0 {
			itemSubtotal = 0
		}
		subtotal += itemSubtotal

		detailList = append(detailList, models.DetailTransaksi{
			ProdukID:    it.ProdukID,
			HargaSatuan: it.HargaSatuan,
			Qty:         it.Qty,
			Subtotal:    itemSubtotal,
		})
	}

	pajakPersen := 11.00
	pajakNominal := math.Round(subtotal * (pajakPersen / 100.0))
	totalBayar := subtotal + pajakNominal

	if req.MetodePembayaran == "" {
		req.MetodePembayaran = "Tunai"
	}

	kembalian := req.UangBayar - totalBayar
	if req.MetodePembayaran == "Tunai" && req.UangBayar > 0 && kembalian < 0 {
		return nil, fmt.Errorf("nominal uang bayar (Rp %.0f) kurang dari total tagihan (Rp %.0f)", req.UangBayar, totalBayar)
	}

	now := time.Now()
	randomNum := rand.Intn(9000) + 1000
	noFaktur := fmt.Sprintf("TRX-%s-%04d", now.Format("20060102"), randomNum)

	trx := &models.Transaksi{
		NomorFaktur:      noFaktur,
		UserID:           req.UserID,
		PelangganID:      req.PelangganID,
		Subtotal:         subtotal,
		PajakPersen:      pajakPersen,
		PajakNominal:     pajakNominal,
		TotalBayar:       totalBayar,
		MetodePembayaran: req.MetodePembayaran,
		StatusTransaksi:  "Selesai",
		CreatedAt:        now,
	}

	if err := s.transaksiRepo.CreateTransactionACID(trx, detailList); err != nil {
		return nil, err
	}

	trx.DetailTransaksi = detailList

	return &CheckoutResponse{
		Transaksi: trx,
		Kembalian: kembalian,
	}, nil
}

func (s *transaksiService) GetAllTransaksi() ([]models.Transaksi, error) {
	return s.transaksiRepo.FindAll()
}

func (s *transaksiService) GetTransaksiByID(id int) (*models.Transaksi, error) {
	return s.transaksiRepo.FindByID(id)
}

func (s *transaksiService) GetTransaksiByNoFaktur(noFaktur string) (*models.Transaksi, error) {
	return s.transaksiRepo.FindByNoFaktur(noFaktur)
}