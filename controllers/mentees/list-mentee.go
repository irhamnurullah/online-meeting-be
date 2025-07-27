package mentees

import (
	"net/http"
	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"
	"time"

	"github.com/gin-gonic/gin"
)

type MenteeResponseDTO struct {
	ProfileID          uint   `json:"id"`
	Fullname           string `json:"fullname"`
	ProfileDescription string `json:"profile_description"`
	Major              string `json:"major"`
	UniversityName     string `json:"university_name"`
	CurrentPosition    string `json:"current_position"`
	ExperienceYear     int    `json:"experience_year"`
}

// GetListMentee godoc
// @Summary List Mentee
// @Description Menampilkan daftar mentee
// @Tags Mentee
// @Security BearerAuth
// @Produce json
// @Router /list-mentee [get]
func GetListMentee(c *gin.Context) {
	db := config.DB

	var profiles []models.Profile
	if err := db.
		Where("is_mentor = ?", false).
		Preload("Educations").
		Preload("Experiences").
		Find(&profiles).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to get mentors", err.Error())
		return
	}

	// Convert ke response DTO
	var mentorDTOs []MenteeResponseDTO
	for _, profile := range profiles {
		var major, universityName, currentPosition string
		var experienceYear int

		if len(profile.Educations) > 0 {
			major = profile.Educations[0].Major
			universityName = profile.Educations[0].UniversityName
		}

		if len(profile.Experiences) > 0 {
			currentPosition = profile.Experiences[0].CurrentPosition
			start := profile.Experiences[0].StartYear
			experienceYear = int(time.Since(start).Hours() / (24 * 365))
		}

		mentorDTOs = append(mentorDTOs, MenteeResponseDTO{
			ProfileID:          profile.ID,
			Fullname:           profile.Fullname,
			ProfileDescription: profile.ProfileDescription,
			Major:              major,
			UniversityName:     universityName,
			CurrentPosition:    currentPosition,
			ExperienceYear:     experienceYear,
		})
	}

	helpers.SuccessResponseWithData(c, http.StatusOK, "Success get mentors", mentorDTOs)
}
