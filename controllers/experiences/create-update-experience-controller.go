package experiences

import (
	"net/http"
	"time"

	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"

	"github.com/gin-gonic/gin"
)

type CreateExperienceRequest struct {
	CurrentPosition string    `json:"current_position"`
	CompanyName     string    `json:"company_name"`
	StartYear       time.Time `json:"start_year"`
	Skills          string    `json:"skills"`
	Achievement     string    `json:"achievement"`
}

// @Summary Create Experience
// @Description Membuat experience baru
// @Tags Experience
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body CreateExperienceRequest true "Experience Data"
// @Router /experience [post]
func CreateExperience(c *gin.Context) {
	db := config.DB

	userIDInterface, exists := c.Get("user_id")
	if !exists {
		helpers.ErrorResponse(c, http.StatusUnauthorized, "unauthorized", "")
		return
	}
	UserID := userIDInterface.(uint)

	var input CreateExperienceRequest
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}

	var existingExperience models.Experience
	err := db.Where("user_id = ?", UserID).First(&existingExperience).Error

	if err == nil {
		// ✅ Sudah ada experience → lakukan update
		existingExperience.CurrentPosition = input.CurrentPosition
		existingExperience.CompanyName = input.CompanyName
		existingExperience.StartYear = input.StartYear
		existingExperience.Skills = input.Skills
		existingExperience.Achievement = input.Achievement

		if err := db.Save(&existingExperience).Error; err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to update experience", err.Error())
			return
		}

		helpers.SuccessResponse(c, http.StatusOK, "experience updated successfully")
		return
	}

	if err.Error() == "record not found" {
		// ✅ Belum ada experience → buat baru
		experience := models.Experience{
			CurrentPosition: input.CurrentPosition,
			StartYear:       input.StartYear,
			CompanyName:     input.CompanyName,
			Skills:          input.Skills,
			Achievement:     input.Achievement,
			UserID:          UserID,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		if err := db.Create(&experience).Error; err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to create experience", err.Error())
			return
		}

		helpers.SuccessResponse(c, http.StatusCreated, "experience created successfully")
		return
	}

	// ❌ Error lain saat mencari experience
	helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to check existing experience", err.Error())
}
