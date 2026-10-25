package handler

import (
	"net/http"
	"pos_material/models"
	"pos_material/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RenderUserPage(c *gin.Context) {
	c.HTML(http.StatusOK, "user.html", nil)
}

func (h *UserHandler) ApiGetUsers(c *gin.Context) {
	var users []models.User
	if err := h.service.GetAllUsers(&users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}