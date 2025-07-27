package routes

import (
	"online-meeting/controllers/appointments"
	"online-meeting/controllers/auth"
	"online-meeting/controllers/educations"
	"online-meeting/controllers/experiences"
	"online-meeting/controllers/mentees"
	"online-meeting/controllers/mentors"
	"online-meeting/controllers/profiles"
	"online-meeting/controllers/rooms"
	"online-meeting/controllers/schedules"
	handler "online-meeting/handlers"
	"online-meeting/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(middlewares.CustomCORS())

	r.POST("/auth/login", auth.AuthLogin)
	r.POST("/auth/register", auth.AuthRegister)

	protected := r.Group("")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.POST("/profile", profiles.CreateProfile)
		protected.GET("/profile", profiles.GetProfile)

		protected.POST("/education", educations.CreateEducation)
		protected.GET("/education", educations.Geteducation)

		protected.POST("/experience", experiences.CreateExperience)
		protected.GET("/experience", experiences.GetExperience)

		protected.GET("/list-mentor", mentors.GetListMentor)
		protected.GET("/mentor/:id", mentors.GetDetailMentor)

		protected.GET("/list-mentee", mentees.GetListMentee)
		protected.GET("/mentee/:id", mentees.GetDetailMentee)

		protected.POST("/appointment/create", appointments.CreateAppointment)
		protected.GET("/appointment/list", appointments.GetAppointments)
		protected.POST("/appointment/update-status/:id", appointments.UpdateStatusAppointment)

		protected.GET("/schedule/list", schedules.GetSchedules)
		protected.POST("/schedule/create", schedules.CreateSchedule)
		protected.PUT("/schedule/update", schedules.UpdateSchedule)
		protected.DELETE("/schedule/delete", schedules.DeleteSchedule)

		protected.POST("/room/create", rooms.CreateRoom)
	}

	r.GET("/ws", func(c *gin.Context) {
		handler.WebSocketHandler(c.Writer, c.Request)
	})
}
