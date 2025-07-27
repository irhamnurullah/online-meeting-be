package schedules

import (
	"net/http"
	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateScheduleRequest struct {
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	AppointmentID uint      `json:"appointment_id"`
}

// CreateSchedule godoc
// @Summary Buat Jadwal
// @Description Menambahkan jadwal ketersediaan mentor
// @Tags Schedule
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body CreateScheduleRequest true "Schedule Data"
// @Router /schedule/create [post]
func CreateSchedule(c *gin.Context) {

	db := config.DB

	var req CreateScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	var appointment models.Appointment
	if err := db.Where("id = ?", req.AppointmentID).First(&appointment).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusNotFound, "Appointment not found", err.Error())
		return
	}

	schedule := models.ScheduleAppointment{
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
		AppointmentID: req.AppointmentID,
	}

	if err := db.Create(&schedule).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to create schedule", err.Error())
	}

	helpers.SuccessResponseWithData(c, http.StatusCreated, "Schedule created successfully", nil)
}
