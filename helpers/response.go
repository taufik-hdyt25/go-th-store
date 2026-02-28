package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JSON(c *gin.Context, status int, message string, data interface{}) {
	res := Response{
		Success: status < 400,
		Status:  status,
		Message: message,
	}
	if data != nil {
		res.Data = data
	}

	c.JSON(status, res)
}

// shortcut helpers (optional tapi enak dipakai)
func Success(c *gin.Context, message string, key string, value interface{}) {
	response := gin.H{
		"success": true,
		"status":  http.StatusOK,
		"message": message,
	}
	if key != "" && value != nil {
		response[key] = value
	}
	c.JSON(http.StatusOK, response)
}

func Created(c *gin.Context, message string, data interface{}) {
	JSON(c, http.StatusCreated, message, data)
}
func Error(c *gin.Context, status int, message string) {
	JSON(c, status, message, nil)
}