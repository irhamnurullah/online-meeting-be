package auth

import (
	"net/http"
	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

// AuthLogin godoc
// @Summary Login
// @Description Login dan mendapatkan JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body LoginRequest true "Login Data"
// @Router /auth/login [post]
func AuthLogin(c *gin.Context) {

	var input LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ErrorResponse(c, http.StatusBadRequest, "Invalid Input", "")
		return
	}

	// cek email
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		helpers.ErrorResponse(c, http.StatusUnauthorized, "User not found", "")
		return
	}

	// verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		helpers.ErrorResponse(c, http.StatusUnauthorized, "Invalid email or password", "")
		return
	}

	token, err := helpers.GenerateToken(user.ID, user.Email)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, "Failed generate token", "")
		return
	}

	loginData := LoginResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: token,
	}

	helpers.SuccessResponseWithData(c, http.StatusCreated, "Success Login", loginData)
}
