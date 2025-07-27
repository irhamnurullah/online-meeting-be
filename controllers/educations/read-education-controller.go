package educations

import (
	"net/http"
	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"

	"github.com/gin-gonic/gin"
)

// GetEducation godoc
// @Summary Ambil Education
// @Description Mengambil data education user yang login
// @Tags Education
// @Security BearerAuth
// @Produce json
// @Router /education [get]
func Geteducation(c *gin.Context) {
	db := config.DB

	userIDInterface, exists := c.Get("user_id")
	if !exists {
		helpers.ErrorResponse(c, http.StatusUnauthorized, "unauthorized", "")
		return
	}

	UserID := userIDInterface.(uint)

	var education models.Education

	if err := db.Where("user_id = ?", UserID).First(&education).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusNotFound, "education not found", err.Error())
		return
	}

	helpers.SuccessResponseWithData(c, http.StatusOK, "Success get education", education)
}
