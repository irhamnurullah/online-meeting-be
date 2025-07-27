package auth

import (
	"net/http"
	"online-meeting/config"
	"online-meeting/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// AuthRegister godoc
// @Summary Register
// @Description Registrasi akun baru
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body RegisterInput true "Register Data"
// @Router /auth/register [post]
func AuthRegister(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data input",
			// "details": err.Error(),
		})

		return
	}

	// cek email exist ?
	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email already taken"})
		return
	}

	// Hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mengenkripsi password"})
		return
	}

	// simpan user
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal mendaftar user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success registered",
		"data": gin.H{
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
