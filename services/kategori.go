package services

import (
	"errors"
	"strings"

	"pos_material/models"
	"pos_material/repositories"
)

type KategoriService interface {
	GetAllKategori() ([]models.Kategori, error)
	GetKategoriByID(id int) (*models.Kategori, error)
	CreateKategori(nama string) (*models.Kategori, error)
	UpdateKategori(id int, nama string) (*models.Kategori, error)
	DeleteKategori(id int) error
}

type kategoriService struct {
	repo repositories.KategoriRepository
}

func NewKategoriService(repo repositories.KategoriRepository) KategoriService {
	return &kategoriService{repo: repo}
}

func (s *kategoriService) GetAllKategori() ([]models.Kategori, error) {
	return s.repo.FindAll()
}

func (s *kategoriService) GetKategoriByID(id int) (*models.Kategori, error) {
	return s.repo.FindByID(id)
}

func (s *kategoriService) CreateKategori(nama string) (*models.Kategori, error) {
	nama = strings.TrimSpace(nama)
	if nama == "" {
		return nil, errors.New("nama kategori tidak boleh kosong")
	}

	kat := &models.Kategori{NamaKategori: nama}
	if err := s.repo.Create(kat); err != nil {
		return nil, err
	}
	return kat, nil
}

func (s *kategoriService) UpdateKategori(id int, nama string) (*models.Kategori, error) {
	nama = strings.TrimSpace(nama)
	if nama == "" {
		return nil, errors.New("nama kategori tidak boleh kosong")
	}

	kat, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	kat.NamaKategori = nama
	if err := s.repo.Update(kat); err != nil {
		return nil, err
	}
	return kat, nil
}

func (s *kategoriService) DeleteKategori(id int) error {
	return s.repo.Delete(id)
}