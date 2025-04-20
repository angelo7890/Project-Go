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
		routes.POST("/events", handler.CreateEventsOpeningHandler)
		routes.GET("/events", handler.ListAllEventsOpeningHandler)

		//endpoints de ingressos
		router.POST("/events/:id/buy", handler.BuyTicketsOpeningHandler)
		router.GET("/tickets/:event_name", handler.ListAllTicketsForEventsOpeningHandler)
		router.GET("/tickets/findAll", handler.ListAllTicketsOpeningHanlder)

	}
}
