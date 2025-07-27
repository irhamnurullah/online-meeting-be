package mentees

import (
	"net/http"
	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetDetailMentee godoc
// @Summary Detail Mentee
// @Description Menampilkan detail mentee berdasarkan ID
// @Tags Mentee
// @Security BearerAuth
// @Produce json
// @Param id path int true "ID Mentee"
// @Router /mentee/{id} [get]
func GetDetailMentee(c *gin.Context) {
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
