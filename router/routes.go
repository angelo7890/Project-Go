package router

import (
	"ingressos-api/handler"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	BASE_PATH := "/api"

	routes := router.Group(BASE_PATH)
	{
		//eventos
		routes.POST("/events", handler.CreateEventsHandler)
		routes.GET("/events", handler.GetAllEventsHandler)
		routes.GET("/events/:id", handler.GetEventForIdHandler)
		//endpoint de delete

		//setor
		routes.POST("/sector", handler.CreateSectorHandler)
		routes.GET("/sector/:id", handler.GetSectorByEventIdHandler)
		routes.DELETE("/sector/:id", handler.DeleteSectorHandler)

		//user
		routes.POST("/user", handler.CreateUserHandler)
		routes.GET("/users", handler.GetAllUsersHandler)
		routes.GET("/user/:id", handler.GetUserByIdHandler)
		//endpoint de delete

		//ingressos
		router.POST("/ticket", handler.BuyTicketsHandler)
		router.GET("/tickets", handler.GetAllTicketsSoldHandler)
		router.GET("/tickets/:id", handler.GetAllTicketsSoldForEventIdHandler)

	}
}
