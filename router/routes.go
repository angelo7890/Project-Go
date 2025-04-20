package router

import "github.com/gin-gonic/gin"

func InitializeRoutes(router *gin.Engine) {
	BASE_PATH := "/api"

	routes := router.Group(BASE_PATH)
	{
		routes.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
