package auth

import (
	"net/http"
	"online-meeting/config"
	"online-meeting/helpers"
	"online-meeting/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func AuthLogin(c *gin.Context) {

	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})

		// err.Error(),
		return
	}

	// cek email
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})

		return
	}

	// verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "wrong email or password",
		})

		return
	}

	token, err := helpers.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed generate token",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Login",
		"data": gin.H{
			"id":           user.ID,
			"name":         user.Name,
			"email":        user.Email,
			"access_token": token,
		},
	})

}
