package rooms

import (
	"net/http"
	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateRoomRequest struct {
	RoomName   string `json:"room_name" binding:"required"`
	ScheduleID uint   `json:"schedule_id" binding:"required"`
}

// CreateRoom godoc
// @Summary Buat Room Meeting
// @Description Membuat room meeting antara mentee dan mentor
// @Tags Room
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body CreateRoomRequest true "Room Data"
// @Router /room/create [post]
func CreateRoom(c *gin.Context) {
	db := config.DB

	var req CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	// Validasi apakah ScheduleID valid
	var schedule models.ScheduleAppointment
	if err := db.First(&schedule, req.ScheduleID).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusNotFound, "Schedule not found", err.Error())
		return
	}

	roomCode := time.Now().Format("20060102150405")

	room := models.Rooms{
		Code:       roomCode,
		RoomName:   req.RoomName,
		ScheduleID: req.ScheduleID,
	}

	if err := db.Create(&room).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to create room", err.Error())
		return
	}

	helpers.SuccessResponseWithData(c, http.StatusCreated, "Room created successfully", gin.H{
		"room_code": roomCode,
		"link":      "/room/" + roomCode,
	})
}
