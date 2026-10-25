package services

import (
	"pos_material/models"
	"pos_material/repositories"
)

type KategoriService struct {
	repo *repositories.KategoriRepository
}

func NewKategoriService(repo *repositories.KategoriRepository) *KategoriService {
	return &KategoriService{repo: repo}
}

func (s *KategoriService) GetAllKategori(kategoris *[]models.Kategori) error {
	return s.repo.FindAll(kategoris)
}

func (s *KategoriService) CreateKategori(kategori *models.Kategori) error {
	return s.repo.Create(kategori)
}