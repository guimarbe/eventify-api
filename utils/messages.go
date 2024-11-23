package utils

import "github.com/gin-gonic/gin"

func RespondWithError(context *gin.Context, statusCode int, message string) {
	context.JSON(statusCode, gin.H{"message": message})
}

func RespondWithMessage(context *gin.Context, statusCode int, message string) {
	context.JSON(statusCode, gin.H{"message": message})
}