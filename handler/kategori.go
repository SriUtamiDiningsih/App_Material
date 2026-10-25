package handler

import (
	"net/http"
	"pos_material/models"
	"pos_material/services"

	"github.com/gin-gonic/gin"
)

type KategoriHandler struct {
	service *services.KategoriService
}

func NewKategoriHandler(service *services.KategoriService) *KategoriHandler {
	return &KategoriHandler{service: service}
}

func (h *KategoriHandler) ApiCreateKategori(c *gin.Context) {
	var kat models.Kategori
	if err := c.ShouldBindJSON(&kat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateKategori(&kat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Kategori berhasil disimpan", "data": kat})
}