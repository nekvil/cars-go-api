package handler

import (
	"github.com/gin-gonic/gin"
)

type successResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func errorResponseJSON(c *gin.Context, status int, err error) {
	c.JSON(status, errorResponse{Error: err.Error()})
}

func successResponseJSON(c *gin.Context, status int, message string, data interface{}) {
	response := successResponse{
		Message: message,
		Data:    data,
	}
	c.JSON(status, response)
}
