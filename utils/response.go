package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response structure for all API responses
type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success response helper
func Success(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// Error response helper
func Error(c *gin.Context, code int, message string) {
	status := "error"
	if code >= 400 && code < 500 {
		status = "fail"
	}
	c.JSON(code, Response{
		Code:    code,
		Status:  status,
		Message: message,
	})
}
