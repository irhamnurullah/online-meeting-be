package appointments

import (
	"net/http"
	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"

	"github.com/gin-gonic/gin"
)

type AppointmentsResponseDTO struct {
	ID        uint                     `json:"id"`
	Objective string                   `json:"objective"`
	Metric    string                   `json:"metric"`
	Chellenge string                   `json:"chellenge"`
	Status    models.AppointmentStatus `json:"status"`
	IDMentor  uint                     `json:"id_mentor"`
	IDMentee  uint                     `json:"id_mentee"`
}

// GetAppointments godoc
// @Summary List Appointment
// @Description Menampilkan semua appointment user
// @Tags Appointment
// @Security BearerAuth
// @Produce json
// @Router /appointment/list [get]
func GetAppointments(c *gin.Context) {
	db := config.DB

	userID, ok := c.Get("user_id")

	if !ok {
		helpers.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", "User not found in context")
		return
	}

	var profile models.Profile
	if err := db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusNotFound, "Profile not found", err.Error())
		return
	}

	var appointments []models.Appointment
	if err := db.
		Where("id_mentor = ? OR id_mentee = ?", profile.ID, profile.ID).
		Find(&appointments).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to get appointments", err.Error())
		return
	}

	// Mapping ke DTO
	var response []AppointmentsResponseDTO
	for _, a := range appointments {
		response = append(response, AppointmentsResponseDTO{
			ID:        a.ID,
			Objective: a.Objective,
			Metric:    a.Metric,
			Chellenge: a.Chellenge,
			Status:    a.Status,
			IDMentor:  a.IDMentor,
			IDMentee:  a.IDMentee,
		})
	}

	helpers.SuccessResponseWithData(c, http.StatusOK, "Success get appointments", response)
}
