package schedules

import (
	"net/http"
	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"

	"github.com/gin-gonic/gin"
)

type DeleteScheduleRequest struct {
	ID uint `json:"id"`
}

// DeleteSchedule godoc
// @Summary Hapus Jadwal
// @Description Menghapus jadwal mentor
// @Tags Schedule
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body DeleteScheduleRequest true "Delete Payload"
// @Router /schedule/delete [delete]
func DeleteSchedule(c *gin.Context) {
	db := config.DB

	var req DeleteScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	var schedule models.ScheduleAppointment
	if err := db.First(&schedule, req.ID).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusNotFound, "Schedule not found", err.Error())
		return
	}

	if err := db.Delete(&schedule).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete schedule", err.Error())
		return
	}

	helpers.SuccessResponse(c, http.StatusOK, "Schedule deleted successfully")
}
