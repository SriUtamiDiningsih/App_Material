package handler

import (
	"net/http"
	"strconv"

	"pos_material/models"
	"pos_material/services"

	"github.com/gin-gonic/gin"
)

type ProdukHandler struct {
	service services.ProdukService
}

func NewProdukHandler(service services.ProdukService) *ProdukHandler {
	return &ProdukHandler{service: service}
}

func (h *ProdukHandler) GetAll(c *gin.Context) {
	search := c.Query("search")
	var kategoriID *int

	if katStr := c.Query("kategori_id"); katStr != "" {
		if id, err := strconv.Atoi(katStr); err == nil {
			kategoriID = &id
		}
	}

	list, err := h.service.GetAllProduk(search, kategoriID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (h *ProdukHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID produk tidak valid"})
		return
	}

	p, err := h.service.GetProdukByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": p})
}

func (h *ProdukHandler) GetByBarcode(c *gin.Context) {
	barcode := c.Param("barcode")
	p, err := h.service.GetProdukByBarcode(barcode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produk barcode tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": p})
}

func (h *ProdukHandler) Create(c *gin.Context) {
	var p models.Produk
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateProduk(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Produk berhasil ditambahkan", "data": p})
}

func (h *ProdukHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID produk tidak valid"})
		return
	}

	var p models.Produk
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateProduk(id, &p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil diperbarui", "data": p})
}

func (h *ProdukHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID produk tidak valid"})
		return
	}

	if err := h.service.DeleteProduk(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
}

func (h *ProdukHandler) GetStats(c *gin.Context) {
	stats, err := h.service.GetProductStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": stats})
}