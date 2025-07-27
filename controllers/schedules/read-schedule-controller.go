package schedules

import (
	"net/http"
	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"

	"github.com/gin-gonic/gin"
)

// GetSchedules godoc
// @Summary List Jadwal
// @Description Menampilkan semua jadwal yang dibuat
// @Tags Schedule
// @Security BearerAuth
// @Produce json
// @Router /schedule/list [get]
func GetSchedules(c *gin.Context) {
	db := config.DB

	userID, ok := c.Get("user_id")
	if !ok {
		helpers.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", "User not found in context")
		return
	}

	// Ambil Profile ID dari user_id
	var profile models.Profile
	if err := db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusNotFound, "Profile not found", err.Error())
		return
	}

	// Ambil list schedule yang berkaitan dengan profile
	var schedules []models.ScheduleAppointment
	if err := db.
		Joins("JOIN appointments ON appointments.id = schedule_appointments.appointment_id").
		Where("appointments.id_mentor = ? OR appointments.id_mentee = ?", profile.ID, profile.ID).
		Preload("Appointment").
		Find(&schedules).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to get schedules", err.Error())
		return
	}

	helpers.SuccessResponseWithData(c, http.StatusOK, "Success get schedules", schedules)
}
