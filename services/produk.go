package services

import (
	"pos_material/models"
	"pos_material/repositories"
)

type ProdukService struct {
	repo *repositories.ProdukRepository
}

func NewProdukService(repo *repositories.ProdukRepository) *ProdukService {
	return &ProdukService{repo: repo}
}

func (s *ProdukService) GetAllProduk(produks *[]models.Produk) error {
	return s.repo.FindAll(produks)
}

func (s *ProdukService) CreateProduk(produk *models.Produk) error {
	return s.repo.Create(produk)
}