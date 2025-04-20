package router

import (
	"ingressos-api/handler"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	BASE_PATH := "/api"

	routes := router.Group(BASE_PATH)
	{
		//endpoints de eventos
		routes.POST("/events", handler.CreateEventsHandler)
		routes.GET("/events", handler.ListAllEventsHandler)

		//endpoints de ingressos
		router.POST("/events/:id/buy", handler.BuyTicketsHandler)
		router.GET("/tickets/:event_name", handler.ListAllTicketsForEventsHandler)
		router.GET("/tickets/findAll", handler.ListAllTicketsHanlder)

	}
}
