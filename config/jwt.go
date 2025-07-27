package config

import "os"

var SecretKey []byte

func InitConfigJwt() {
	SecretKey = []byte(os.Getenv("JWT_SECRET"))
}

func JwtSecret() []byte {
	return SecretKey
}
