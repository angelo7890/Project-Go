package responses

import "github.com/gin-gonic/gin"

func SendError(context *gin.Context, code int, msg string) {
	context.Header("Content-type", "application/json")
	context.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}
