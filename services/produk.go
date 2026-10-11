package services

import (
	"errors"
	"strings"

	"pos_material/models"
	"pos_material/repositories"
)

type ProdukService interface {
	GetAllProduk(search string, kategoriID *int) ([]models.Produk, error)
	GetProdukByID(id int) (*models.Produk, error)
	GetProdukByBarcode(barcode string) (*models.Produk, error)
	CreateProduk(p *models.Produk) error
	UpdateProduk(id int, p *models.Produk) error
	DeleteProduk(id int) error
	GetProductStats() (map[string]int64, error)
}

type produkService struct {
	repo repositories.ProdukRepository
}

func NewProdukService(repo repositories.ProdukRepository) ProdukService {
	return &produkService{repo: repo}
}

func (s *produkService) GetAllProduk(search string, kategoriID *int) ([]models.Produk, error) {
	return s.repo.FindAll(search, kategoriID)
}

func (s *produkService) GetProdukByID(id int) (*models.Produk, error) {
	return s.repo.FindByID(id)
}

func (s *produkService) GetProdukByBarcode(barcode string) (*models.Produk, error) {
	barcode = strings.TrimSpace(barcode)
	if barcode == "" {
		return nil, errors.New("barcode tidak boleh kosong")
	}
	return s.repo.FindByBarcode(barcode)
}

func (s *produkService) CreateProduk(p *models.Produk) error {
	p.NamaProduk = strings.TrimSpace(p.NamaProduk)
	if p.NamaProduk == "" {
		return errors.New("nama produk wajib diisi")
	}
	if p.Satuan == "" {
		return errors.New("satuan produk wajib diisi")
	}
	if p.HargaJual < 0 {
		return errors.New("harga jual tidak boleh negatif")
	}
	if p.Stok < 0 {
		p.Stok = 0
	}
	return s.repo.Create(p)
}

func (s *produkService) UpdateProduk(id int, p *models.Produk) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	p.NamaProduk = strings.TrimSpace(p.NamaProduk)
	if p.NamaProduk == "" {
		return errors.New("nama produk wajib diisi")
	}

	existing.NamaProduk = p.NamaProduk
	existing.KategoriID = p.KategoriID
	existing.Barcode = p.Barcode
	existing.Satuan = p.Satuan
	existing.HargaJual = p.HargaJual
	existing.Stok = p.Stok
	existing.GambarURL = p.GambarURL

	return s.repo.Update(existing)
}

func (s *produkService) DeleteProduk(id int) error {
	return s.repo.Delete(id)
}

func (s *produkService) GetProductStats() (map[string]int64, error) {
	total, err := s.repo.CountTotal()
	if err != nil {
		return nil, err
	}
	inStock, err := s.repo.CountInStock()
	if err != nil {
		return nil, err
	}
	lowStock, err := s.repo.CountLowStock(25)
	if err != nil {
		return nil, err
	}

	return map[string]int64{
		"total_produk":  total,
		"stok_tersedia": inStock,
		"stok_rendah":   lowStock,
	}, nil
}