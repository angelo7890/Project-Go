package handler

import (
	"fmt"
	"ingressos-api/dto"
	"ingressos-api/repository"
	responses "ingressos-api/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BuyTicketsHandler(context *gin.Context) {
	var request dto.BuyTicketRequestDTO
	if err := context.ShouldBindJSON(&request); err != nil {
		responses.SendError(context, http.StatusBadRequest, fmt.Sprintf("Erro ao processar os dados: %v", err))
		return
	}

	// Validação dos dados
	if err := request.Validade(); err != nil {
		responses.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	// Abrir a transação
	tx, err := db.Begin()
	if err != nil {
		responses.SendError(context, http.StatusInternalServerError, fmt.Sprintf("Erro ao iniciar transação: %v", err))
		return
	}
	defer tx.Rollback() // Se houver erro, faz o rollback

	// Chama a função de compra de ingresso
	ticket, err := repository.BuyTicket(tx, request, db)
	if err != nil {
		responses.SendError(context, http.StatusInternalServerError, fmt.Sprintf("Erro ao comprar ingresso: %v", err))
		return
	}

	// Confirma a transação
	if err := tx.Commit(); err != nil {
		responses.SendError(context, http.StatusInternalServerError, fmt.Sprintf("Erro ao confirmar a transação: %v", err))
		return
	}
	responses.SendSuccess(context, "Compra de ingresso", ticket)
}

func GetAllTicketsSoldHandler(context *gin.Context) {
}

func GetAllTicketsSoldForEventIdHandler(context *gin.Context) {
}
