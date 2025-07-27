package educations

import (
	"net/http"
	"time"

	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"

	"github.com/gin-gonic/gin"
)

type CreateEducationRequest struct {
	UniversityName string    `json:"university_name"`
	Major          string    `json:"major"`
	EndYear        time.Time `json:"end_year"`
}

// @Summary Create Education
// @Description Membuat education baru
// @Tags Education
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body CreateEducationRequest true "Education Data"
// @Router /education [post]
func CreateEducation(c *gin.Context) {
	db := config.DB

	userIDInterface, exists := c.Get("user_id")
	if !exists {
		helpers.ErrorResponse(c, http.StatusUnauthorized, "unauthorized", "")
		return
	}
	UserID := userIDInterface.(uint)

	var input CreateEducationRequest
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}

	var existingEducation models.Education
	err := db.Where("user_id = ?", UserID).First(&existingEducation).Error

	if err == nil {
		// ✅ Sudah ada education → lakukan update
		existingEducation.UniversityName = input.UniversityName
		existingEducation.EndYear = input.EndYear
		existingEducation.Major = input.Major
		existingEducation.UpdatedAt = time.Now()

		if err := db.Save(&existingEducation).Error; err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to update education", err.Error())
			return
		}

		helpers.SuccessResponse(c, http.StatusOK, "education updated successfully")
		return
	}

	if err.Error() == "record not found" {
		// ✅ Belum ada education → buat baru
		education := models.Education{
			UniversityName: input.UniversityName,
			EndYear:        input.EndYear,
			Major:          input.Major,
			UserID:         UserID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if err := db.Create(&education).Error; err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to create education", err.Error())
			return
		}

		helpers.SuccessResponse(c, http.StatusCreated, "education created successfully")
		return
	}

	// ❌ Error lain saat mencari education
	helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to check existing education", err.Error())
}
