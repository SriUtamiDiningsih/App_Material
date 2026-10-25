package handler

import (
	"net/http"
	"pos_material/models"
	"pos_material/services"

	"github.com/gin-gonic/gin"
)

type TransaksiHandler struct {
	service *services.TransaksiService
}

func NewTransaksiHandler(service *services.TransaksiService) *TransaksiHandler {
	return &TransaksiHandler{service: service}
}

func (h *TransaksiHandler) RenderLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func (h *TransaksiHandler) RenderTransaksiPage(c *gin.Context) {
	c.HTML(http.StatusOK, "transaksi.html", nil)
}

func (h *TransaksiHandler) RenderReportTransaksi(c *gin.Context) {
	c.HTML(http.StatusOK, "laporan_transaksi.html", nil)
}

func (h *TransaksiHandler) ApiGetTransaksi(c *gin.Context) {
	var trans []models.Transaksi
	if err := h.service.GetAllTransaksis(&trans); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trans)
}

func (h *TransaksiHandler) ApiGetDetailPenjualan(c *gin.Context) {
	var details []models.DetailTransaksi
	if err := h.service.GetAllDetails(&details); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, details)
}

func (h *TransaksiHandler) ApiCheckout(c *gin.Context) {
	var trx models.Transaksi
	if err := c.ShouldBindJSON(&trx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.ProcessCheckout(&trx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Transaksi Sukses! Stok produk berkurang, detail tersimpan, dan struk siap dicetak.",
		"data":    trx,
	})
}