package handler

import (
	"net/http"
	"strconv"

	"pos_material/services"

	"github.com/gin-gonic/gin"
)

type KategoriHandler struct {
	service services.KategoriService
}

func NewKategoriHandler(service services.KategoriService) *KategoriHandler {
	return &KategoriHandler{service: service}
}

func (h *KategoriHandler) GetAll(c *gin.Context) {
	list, err := h.service.GetAllKategori()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (h *KategoriHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID kategori tidak valid"})
		return
	}

	kat, err := h.service.GetKategoriByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": kat})
}

type KategoriInput struct {
	NamaKategori string `json:"nama_kategori" binding:"required"`
}

func (h *KategoriHandler) Create(c *gin.Context) {
	var input KategoriInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kat, err := h.service.CreateKategori(input.NamaKategori)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Kategori berhasil ditambahkan", "data": kat})
}

func (h *KategoriHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var input KategoriInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kat, err := h.service.UpdateKategori(id, input.NamaKategori)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Kategori berhasil diperbarui", "data": kat})
}

func (h *KategoriHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	if err := h.service.DeleteKategori(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Kategori berhasil dihapus"})
}