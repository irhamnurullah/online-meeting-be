package experiences

import (
	"net/http"
	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"

	"github.com/gin-gonic/gin"
)

// GetExperience godoc
// @Summary Ambil Experience
// @Description Mengambil data Experience user yang login
// @Tags Experience
// @Security BearerAuth
// @Produce json
// @Router /experience [get]
func GetExperience(c *gin.Context) {
	db := config.DB

	userIDInterface, exists := c.Get("user_id")
	if !exists {
		helpers.ErrorResponse(c, http.StatusUnauthorized, "unauthorized", "")
		return
	}

	UserID := userIDInterface.(uint)

	var experience models.Experience

	if err := db.Where("user_id = ?", UserID).First(&experience).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusNotFound, "Experience not found", err.Error())
		return
	}

	helpers.SuccessResponseWithData(c, http.StatusOK, "Success get Experience", experience)
}
