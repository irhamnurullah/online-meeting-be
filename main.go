package main

import (
	"online-meeting/config"
	"online-meeting/controllers"
	"online-meeting/controllers/auth"
	"online-meeting/controllers/rooms"
	handler "online-meeting/handlers"
	"online-meeting/middlewares"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectionDatabase()

	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // ganti sesuai asal FE kamu
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/auth/login", auth.AuthLogin)
	r.POST("/auth/register", auth.AuthRegister)

	protected := r.Group("")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/users", controllers.GetUsers)
		protected.POST("/create-room", rooms.CreateRoom)

	}
	r.GET("/ws", func(c *gin.Context) {
		// roomCode := c.Query("room_code")
		// if roomCode == "" {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing room_code"})
		// 	return
		// }

		// userID := c.MustGet("user_id").(string) // Diambil dari AuthMiddleware
		handler.WebSocketHandler(c.Writer, c.Request)
	})

	// r.GET("/seed", func(c *gin.Context) {
	// 	config.SeedData()
	// 	c.JSON(200, gin.H{"message": "Seeding selesai"})
	// })

	r.Run(":8080")
}
