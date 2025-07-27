package helpers

import (
	"errors"
	"online-meeting/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId uint, email string) (string, error) {
	secret := config.JwtSecret()
	claims := jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

type TokenClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func ParseToken(tokenString string) (*TokenClaims, error) {
	claims := &TokenClaims{}
	secret := config.JwtSecret()

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Optional: cek expired manual juga
	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	return claims, nil
}
