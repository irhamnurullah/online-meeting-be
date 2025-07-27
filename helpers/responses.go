package helpers

import (
	"github.com/gin-gonic/gin"
)

type APIResponses struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type ErrorResponses struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type MessageResponses struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func SuccessResponse(c *gin.Context, statusCode int, message string) {
	response := APIResponses{
		Success: true,
		Message: message,
	}
	c.JSON(statusCode, response)
}

func SuccessResponseWithData(c *gin.Context, statusCode int, message string, data any) {
	response := APIResponses{
		Success: true,
		Message: message,
		Data:    data,
	}
	c.JSON(statusCode, response)
}

// ErrorResponse sends an error response
func ErrorResponse(c *gin.Context, statusCode int, message string, error string) {
	response := ErrorResponses{
		Success: false,
		Message: message,
		Error:   error,
	}
	c.JSON(statusCode, response)
}

// MessageResponse sends a simple message response
func MessageResponse(c *gin.Context, statusCode int, message string) {
	response := MessageResponses{
		Success: true,
		Message: message,
	}
	c.JSON(statusCode, response)
}
