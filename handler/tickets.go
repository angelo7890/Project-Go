package handler

import (
	"fmt"
	"ingressos-api/database"
	"ingressos-api/dto"
	"ingressos-api/repository"
	responses "ingressos-api/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BuyTicketsHandler(context *gin.Context) {
	var request dto.BuyTicketRequestDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		responses.SendError(context, http.StatusBadRequest, fmt.Sprintf("Dados inválidos: %v", err))
		return
	}

	db := database.GetDB()
	transaction, err := db.Begin()
	if err != nil {
		responses.SendError(context, http.StatusInternalServerError, "Erro ao iniciar transação")
		return
	}

	defer func() {
		if err != nil {
			transaction.Rollback()
		}
	}()

	ticket, err := repository.BuyTicket(transaction, request)
	if err != nil {
		responses.SendError(context, http.StatusBadRequest, fmt.Sprintf("Erro ao comprar ingresso: %v", err))
		return
	}

	if err = transaction.Commit(); err != nil {
		responses.SendError(context, http.StatusInternalServerError, fmt.Sprintf("Erro ao confirmar transação: %v", err))
		return
	}

	responses.SendSuccess(context, "Ingresso comprado com sucesso", ticket)
}

func GetAllTicketsSoldHandler(context *gin.Context) {
	db := database.GetDB()

	tickets, err := repository.GetAllTicketsSoldRepository(db)
	if err != nil {
		responses.SendError(context, http.StatusBadRequest, "nao foi possivel buscar ingressos")
	}
	responses.SendSuccess(context, "get-all-tickets-sold", tickets)
}

func GetAllTicketsSoldForEventIdHandler(context *gin.Context) {
	idParam := context.Param("id")
	idEvent, err := strconv.Atoi(idParam)
	if err != nil {
		responses.SendError(context, http.StatusBadRequest, "id invalido")
		return
	}
	db := database.GetDB()
	tickets, err := repository.GetAllTicketsSoldByEventIDRepository(db, idEvent)
	if err != nil {
		responses.SendError(context, http.StatusBadRequest, "nao foi possivel buscar ingressos")
	}
	responses.SendSuccess(context, "get-all-tickets-sold-by-event-id", tickets)
}
