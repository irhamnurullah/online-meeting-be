package appointments

import (
	"net/http"
	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateStatusRequest struct {
	Status models.AppointmentStatus `json:"status" binding:"required"`
}

// UpdateStatusAppointment godoc
// @Summary Update Status Appointment
// @Description Mengubah status appointment (booking/scheduled/done)
// @Tags Appointment
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID Appointment"
// @Param body body UpdateStatusRequest true "Status Update"
// @Router /appointment/update-status/{id} [post]
func UpdateStatusAppointment(c *gin.Context) {
	db := config.DB
	appointmentID := c.Param("id")

	// Validasi dan parsing ID
	id, err := strconv.Atoi(appointmentID)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid appointment ID", err.Error())
		return
	}

	// Bind JSON
	var req UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return
	}

	// Cari appointment
	var appointment models.Appointment
	if err := db.First(&appointment, id).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusNotFound, "Appointment not found", err.Error())
		return
	}

	// Update status
	appointment.Status = req.Status
	if err := db.Save(&appointment).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to update status", err.Error())
		return
	}

	helpers.SuccessResponseWithData(c, http.StatusOK, "Appointment status updated", appointment)
}
