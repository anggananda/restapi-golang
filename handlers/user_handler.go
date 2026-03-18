package handlers

import (
	"context"
	"net/http"
	"restapi-golang/services"
	"restapi-golang/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

// GetDataProfile mendapatkan  detail mahasiswa
// @Summary      Get data profile
// @Description  Mendapatkan data profile
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200           {object}  models.ListDetailResponse{datas=models.UserAuth}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /user/details [get]
func (h *UserHandler) GetDataProfile(c *gin.Context) {
	username, err := utils.GetUsername(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	contentType := c.DefaultQuery("contentType", "json")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	profile, err := h.UserService.CheckUserByUsername(ctx, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if contentType == "msgpack" {
		utils.Render(c, http.StatusOK, gin.H{"datas": profile, "status": "success"})
	} else {
		c.JSON(http.StatusOK, gin.H{"datas": profile, "status": "success"})
	}
}
