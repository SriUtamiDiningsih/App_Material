package handler

import (
	"net/http"
	"pos_material/models"
	"pos_material/services"

	"github.com/gin-gonic/gin"
)

type ProdukHandler struct {
	service *services.ProdukService
}

func NewProdukHandler(service *services.ProdukService) *ProdukHandler {
	return &ProdukHandler{service: service}
}

func (h *ProdukHandler) RenderProdukPage(c *gin.Context) {
	c.HTML(http.StatusOK, "produk.html", nil)
}

func (h *ProdukHandler) RenderReportProduk(c *gin.Context) {
	c.HTML(http.StatusOK, "laporan_produk.html", nil)
}

func (h *ProdukHandler) ApiGetProduk(c *gin.Context) {
	var produks []models.Produk
	if err := h.service.GetAllProduk(&produks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, produks)
}

func (h *ProdukHandler) ApiCreateProduk(c *gin.Context) {
	var produk models.Produk
	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateProduk(&produk); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Produk berhasil disimpan ke database", "data": produk})
}