package services

import (
	"errors"
	"fmt"
	"pos_material/models"
	"pos_material/repositories"
	"time"
)

type TransaksiService struct {
	repo *repositories.TransaksiRepository
}

func NewTransaksiService(repo *repositories.TransaksiRepository) *TransaksiService {
	return &TransaksiService{repo: repo}
}

func (s *TransaksiService) GetAllTransaksis(transaksis *[]models.Transaksi) error {
	return s.repo.FindAllTransaksis(transaksis)
}

func (s *TransaksiService) GetAllDetails(details *[]models.DetailTransaksi) error {
	return s.repo.FindAllDetails(details)
}

func (s *TransaksiService) ProcessCheckout(trx *models.Transaksi) error {
	trx.NomorFaktur = fmt.Sprintf("INV-%s", time.Now().Format("20060102150405"))
	
	var subtotal float64
	for i := range trx.Details {
		item := &trx.Details[i]
		item.Subtotal = item.HargaSatuan * float64(item.Qty)
		subtotal += item.Subtotal
	}

	trx.Subtotal = subtotal
	trx.PajakPersen = 11.00
	trx.PajakNominal = subtotal * 0.11
	trx.TotalBayar = trx.Subtotal + trx.PajakNominal

	if err := s.repo.CreateCheckout(trx); err != nil {
		return errors.New("gagal memproses transaksi: stok tidak cukup atau database error")
	}
	return nil
}