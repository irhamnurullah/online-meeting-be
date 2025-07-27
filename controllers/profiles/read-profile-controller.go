package profiles

import (
	"net/http"
	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"

	"github.com/gin-gonic/gin"
)

// GetProfile godoc
// @Summary Ambil Profile
// @Description Mengambil data profile user yang login
// @Tags Profile
// @Security BearerAuth
// @Produce json
// @Router /profile [get]
func GetProfile(c *gin.Context) {
	db := config.DB

	userIDInterface, exists := c.Get("user_id")
	if !exists {
		helpers.ErrorResponse(c, http.StatusUnauthorized, "unauthorized", "")
		return
	}

	UserID := userIDInterface.(uint)

	var profile models.Profile

	if err := db.Where("user_id = ?", UserID).First(&profile).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusNotFound, "Profile not found", err.Error())
		return
	}

	helpers.SuccessResponseWithData(c, http.StatusOK, "Success get Profile", profile)
}
