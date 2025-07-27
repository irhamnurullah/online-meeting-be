package appointments

import (
	"net/http"

	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"

	"github.com/gin-gonic/gin"
)

type CreateAppointmentRequest struct {
	Objective string                   `json:"objective"`
	Metric    string                   `json:"metric"`
	Chellenge string                   `json:"chellenge"`
	Status    models.AppointmentStatus `json:"status"`
	IDMentor  uint                     `json:"id_mentor"`
	IDMentee  uint                     `json:"id_metee"`
}

// CreateAppointment godoc
// @Summary Buat Janji
// @Description Membuat janji temu antara mentee dan mentor
// @Tags Appointment
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body CreateAppointmentRequest true "Appointment Data"
// @Router /appointment/create [post]
func CreateAppointment(c *gin.Context) {
	db := config.DB

	// Ambil user_id dari JWT (disimpan di middleware auth)
	userID, ok := c.Get("user_id")
	if !ok {
		helpers.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", "User not found in context")
		return
	}

	var req CreateAppointmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	// Cari ID mentee dari table Profile berdasarkan user_id
	var profile models.Profile
	if err := db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusNotFound, "Profile not found", err.Error())
		return
	}

	// Set default status jika tidak dikirim
	if req.Status == "" {
		req.Status = models.StatusBooking
	}

	appointment := models.Appointment{
		Objective: req.Objective,
		Metric:    req.Metric,
		Chellenge: req.Chellenge,
		Status:    req.Status,
		IDMentor:  req.IDMentor,
		IDMentee:  profile.ID, // diambil dari relasi profile mentee
	}

	if err := db.Create(&appointment).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed to create appointment", err.Error())
		return
	}

	helpers.SuccessResponseWithData(c, http.StatusCreated, "Appointment created successfully", nil)
}
