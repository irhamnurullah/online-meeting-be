package mentors

import (
	"net/http"
	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetDetailMentor godoc
// @Summary Detail Mentor
// @Description Menampilkan detail mentor berdasarkan ID
// @Tags Mentor
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID Mentor"
// @Router /mentor/{id} [get]
func GetDetailMentor(c *gin.Context) {
	db := config.DB

	// Ambil param id dari URL
	profileIDParam := c.Param("id")
	profileID, err := strconv.ParseUint(profileIDParam, 10, 64)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid profile ID", err.Error())
		return
	}

	var profile models.Profile
	if err := db.
		Where("id = ?", profileID).
		Preload("Educations").
		Preload("Experiences").
		First(&profile).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusNotFound, "Mentor not found", err.Error())
		return
	}

	helpers.SuccessResponseWithData(c, http.StatusOK, "Success get mentor detail", profile)
}
