package handlers

import (
	"blog-backend/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespondError(c *gin.Context, status int, message string, err error) {
	if err != nil {
		config.Logger.Println(message, err)
	}
	c.JSON(status, gin.H{"error": message})
}

func RespondSuccess(c *gin.Context, message string, data gin.H) {
	config.Logger.Println(message)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
	})
}
