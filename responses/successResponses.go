package responses

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSuccess(context *gin.Context, op string, data interface{}) {
	context.Header("Content-type", "application/json")
	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}
