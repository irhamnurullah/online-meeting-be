package rooms

import (
	"net/http"
	"online-meeting/config"
	"online-meeting/models"

	"time"

	"github.com/gin-gonic/gin"
)

type CreateRoomRequest struct {
	RoomName  string    `json:"room_name" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
}

func CreateRoom(c *gin.Context) {

	db := config.DB

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})

		return
	}

	var req CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	roomCode := time.Now().Format("20060102150405")

	room := models.Room{
		Code:      roomCode,
		RoomName:  req.RoomName,
		StartTime: req.StartTime,
		EndTime:   req.StartTime.Add(1 * time.Hour),
		HostID:    user.(models.User).ID,
	}

	if err := db.Create(&room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create room",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Room created successfully",
		"room_code": roomCode,
	})

}
