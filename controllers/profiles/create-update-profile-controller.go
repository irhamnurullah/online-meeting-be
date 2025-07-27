package profiles

import (
	"net/http"
	"time"

	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"

	"github.com/gin-gonic/gin"
)

type CreateProfileRequest struct {
	Fullname           string    `json:"fullname" binding:"required,min=2,max=100"`
	DateBirth          time.Time `json:"date_birth" binding:"required"`
	ProfileDescription string    `json:"profile_description" binding:"omitempty,max=1000"`
	ContactPerson      string    `json:"contact_person" binding:"omitempty,max=100"`
	ProfilePicture     string    `json:"profile_picture" binding:"omitempty,url"`
	IsMentor           bool      `json:"is_mentor"`
}

// @Summary Create Profile
// @Description Membuat profil baru
// @Tags Profile
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body CreateProfileRequest true "Profile Data"
// @Router /profile [post]
func CreateProfile(c *gin.Context) {
	db := config.DB

	userIDInterface, exists := c.Get("user_id")
	if !exists {
		helpers.ErrorResponse(c, http.StatusUnauthorized, "unauthorized", "")
		return
	}
	UserID := userIDInterface.(uint)

	var input CreateProfileRequest
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}

	var existingProfile models.Profile
	err := db.Where("user_id = ?", UserID).First(&existingProfile).Error

	if err == nil {
		// ✅ Sudah ada profile → lakukan update
		existingProfile.Fullname = input.Fullname
		existingProfile.DateBirth = input.DateBirth
		existingProfile.ProfileDescription = input.ProfileDescription
		existingProfile.ContactPerson = input.ContactPerson
		existingProfile.ProfilePicture = input.ProfilePicture
		existingProfile.IsMentor = input.IsMentor
		existingProfile.UpdatedAt = time.Now()

		if err := db.Save(&existingProfile).Error; err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to update profile", err.Error())
			return
		}

		helpers.SuccessResponse(c, http.StatusOK, "Profile updated successfully")
		return
	}

	if err.Error() == "record not found" {
		// ✅ Belum ada profile → buat baru
		profile := models.Profile{
			Fullname:           input.Fullname,
			DateBirth:          input.DateBirth,
			ProfileDescription: input.ProfileDescription,
			ContactPerson:      input.ContactPerson,
			ProfilePicture:     input.ProfilePicture,
			IsMentor:           input.IsMentor,
			UserID:             UserID,
			CreatedAt:          time.Now(),
			UpdatedAt:          time.Now(),
		}

		if err := db.Create(&profile).Error; err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to create profile", err.Error())
			return
		}

		helpers.SuccessResponse(c, http.StatusCreated, "Profile created successfully")
		return
	}

	// ❌ Error lain saat mencari profile
	helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to check existing profile", err.Error())
}
