package schedules

import (
	"net/http"
	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateScheduleRequest struct {
	ID            uint      `json:"id"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	AppointmentID uint      `json:"appointment_id"`
}

// UpdateSchedule godoc
// @Summary Ubah Jadwal
// @Description Mengubah jadwal ketersediaan mentor
// @Tags Schedule
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body UpdateScheduleRequest true "Update Data"
// @Router /schedule/update [put]
func UpdateSchedule(c *gin.Context) {

	db := config.DB

	var req UpdateScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	var appointment models.Appointment
	if err := db.Where("id = ?", req.AppointmentID).First(&appointment).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusNotFound, "Appointment not found", err.Error())
		return
	}

	var existingSchedule models.ScheduleAppointment
	err := db.Where("id = ?", req.ID).First(&existingSchedule).Error
	if err == nil {
		existingSchedule.StartDate = req.StartDate
		existingSchedule.EndDate = req.EndDate

		if err := db.Save(&existingSchedule).Error; err != nil {
			helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to create schedule", err.Error())
			return
		}

		helpers.SuccessResponseWithData(c, http.StatusCreated, "Schedule updated successfully", existingSchedule)
		return
	}

	helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to check existing schedule", err.Error())

}
