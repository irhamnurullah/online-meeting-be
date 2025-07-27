package main

import (
	"online-meeting/config"
	"online-meeting/routes"

	"github.com/gin-gonic/gin"

	_ "online-meeting/docs" // import ini penting untuk swag membaca file docs.go

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Online Meeting API
// @version 1.0
// @description Ini adalah dokumentasi API untuk aplikasi online meeting.
// @termsOfService http://swagger.io/terms/

// @contact.name Tim Developer
// @contact.email support@example.com

// @BasePath /

// ✅ Tambahkan ini di sini
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config.ConnectionDatabase()
	config.LoadEnv()
	config.InitConfigJwt()

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	routes.SetupRoutes(r)

	port := config.GetEnv("APP_PORT")
	r.Run(":" + port)
}
