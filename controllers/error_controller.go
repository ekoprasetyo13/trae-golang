package controllers

import (
	"net/http"
	"sample-api/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Trigger 400 Bad Request
// @Tags errors
// @Router /errors/400 [get]
func Error400(c *gin.Context) {
	utils.Error(c, http.StatusBadRequest, "This is a 400 Bad Request error")
}

// @Summary Trigger 401 Unauthorized
// @Tags errors
// @Router /errors/401 [get]
func Error401(c *gin.Context) {
	utils.Error(c, http.StatusUnauthorized, "This is a 401 Unauthorized error")
}

// @Summary Trigger 403 Forbidden
// @Tags errors
// @Router /errors/403 [get]
func Error403(c *gin.Context) {
	utils.Error(c, http.StatusForbidden, "This is a 403 Forbidden error")
}

// @Summary Trigger 404 Not Found
// @Tags errors
// @Router /errors/404 [get]
func Error404(c *gin.Context) {
	utils.Error(c, http.StatusNotFound, "This is a 404 Not Found error")
}

// @Summary Trigger 405 Method Not Allowed
// @Tags errors
// @Router /errors/405 [get]
func Error405(c *gin.Context) {
	utils.Error(c, http.StatusMethodNotAllowed, "This is a 405 Method Not Allowed error")
}

// @Summary Trigger 408 Request Timeout
// @Tags errors
// @Router /errors/408 [get]
func Error408(c *gin.Context) {
	utils.Error(c, http.StatusRequestTimeout, "This is a 408 Request Timeout error")
}

// @Summary Trigger 500 Internal Server Error
// @Tags errors
// @Router /errors/500 [get]
func Error500(c *gin.Context) {
	utils.Error(c, http.StatusInternalServerError, "This is a 500 Internal Server Error")
}

// @Summary Trigger 503 Service Unavailable
// @Tags errors
// @Router /errors/503 [get]
func Error503(c *gin.Context) {
	utils.Error(c, http.StatusServiceUnavailable, "This is a 503 Service Unavailable error")
}

// @Summary Trigger 504 Gateway Timeout
// @Tags errors
// @Router /errors/504 [get]
func Error504(c *gin.Context) {
	utils.Error(c, http.StatusGatewayTimeout, "This is a 504 Gateway Timeout error")
}
