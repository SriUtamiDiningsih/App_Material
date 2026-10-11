package handler

import (
	"net/http"
	"strconv"

	"pos_material/services"

	"github.com/gin-gonic/gin"
)

type TransaksiHandler struct {
	service services.TransaksiService
}

func NewTransaksiHandler(service services.TransaksiService) *TransaksiHandler {
	return &TransaksiHandler{service: service}
}

func (h *TransaksiHandler) Checkout(c *gin.Context) {
	var req services.CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.ProcessCheckout(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Transaksi berhasil disimpan dan stok otomatis terpotong!",
		"transaksi": res.Transaksi,
		"kembalian": res.Kembalian,
	})
}

func (h *TransaksiHandler) GetAll(c *gin.Context) {
	list, err := h.service.GetAllTransaksi()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": list})
}

func (h *TransaksiHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID transaksi tidak valid"})
		return
	}

	trx, err := h.service.GetTransaksiByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaksi tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": trx})
}

func (h *TransaksiHandler) GetByNoFaktur(c *gin.Context) {
	noFaktur := c.Param("no_faktur")
	trx, err := h.service.GetTransaksiByNoFaktur(noFaktur)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaksi tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": trx})
}