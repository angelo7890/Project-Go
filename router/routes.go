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
		routes.DELETE("/events/:id", handler.DeleteEventById)

		//setor
		routes.POST("/sector", handler.CreateSectorHandler)
		routes.GET("/sector/:id", handler.GetSectorByEventIdHandler)
		routes.DELETE("/sector/:id", handler.DeleteSectorHandler)
		routes.PATCH("/updateSector", handler.UpdateTicketSectorHandler)

		//user
		routes.POST("/user", handler.CreateUserHandler)
		routes.GET("/users", handler.GetAllUsersHandler)
		routes.GET("/user/:id", handler.GetUserByIdHandler)
		routes.DELETE("/user/:id", handler.DeleteUserById)

		//ingressos
		routes.POST("/ticket", handler.BuyTicketsHandler)
		routes.GET("/tickets", handler.GetAllTicketsSoldHandler)
		routes.GET("/tickets/:id", handler.GetAllTicketsSoldForEventIdHandler)
	}
}
