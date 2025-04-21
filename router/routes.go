package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	BASE_PATH := "/api"

	routes := router.Group(BASE_PATH)
	{
		//eventos
		routes.POST("/events")
		routes.GET("/events")
		routes.GET("/events/{id}")

		//setor
		routes.GET("/events/{id}")
		routes.GET("/events/{id}")
		routes.GET("/events/{id}")

		//user
		routes.GET("/events/{id}")
		routes.GET("/events/{id}")
		routes.GET("/events/{id}")

		//ingressos
		router.POST("/sales")
		router.GET("/tickets/:event_name")
		router.GET("/tickets/findAll")

	}
}
