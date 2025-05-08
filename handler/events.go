package handler

import (
	"ingressos-api/database"
	"ingressos-api/dto"
	"ingressos-api/repository"
	responses "ingressos-api/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateEventsHandler(context *gin.Context) {
	request := dto.EventCreateDTO{}
	db := database.GetDB()

	if err := context.ShouldBindJSON(&request); err != nil {

		responses.SendError(context, http.StatusBadRequest, "Dados inválidos.")
		return
	}

	if err := request.Validate(); err != nil {
		responses.SendError(context, http.StatusBadRequest, "dados invalidos")
		return
	}

	event := dto.EventCreateDTO{
		Name:     request.Name,
		Date:     request.Date,
		Location: request.Location,
	}
	if err := repository.CreateEvent(db, event); err != nil {
		responses.SendError(context, http.StatusInternalServerError, "erro ao salvar evento")
		return
	}
	responses.SendSuccess(context, "create-event", event)
}

func GetAllEventsHandler(context *gin.Context) {
	db := database.GetDB()

	data, err := repository.GetAllEvents(db)
	if err != nil {
		responses.SendError(context, http.StatusBadRequest, "erro ao buscar eventos")
		return
	}
	responses.SendSuccess(context, "get-all-events", data)
}

func GetEventForIdHandler(context *gin.Context) {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		responses.SendError(context, http.StatusBadRequest, "id invalido")
		return
	}
	db := database.GetDB()

	data, err := repository.GetEventByID(db, id)
	if err != nil {
		responses.SendError(context, http.StatusBadRequest, "erro ao buscar evento")
		return
	}
	responses.SendSuccess(context, "get-event-by-id", data)

}
func DeleteEventById(context *gin.Context) {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		responses.SendError(context, http.StatusBadRequest, "ID inválido")
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
	if err = repository.DeleteEventById(transaction, id); err != nil {
		responses.SendError(context, http.StatusBadRequest, "Não foi possível deletar o evento: "+err.Error())
		return
	}

	if err = transaction.Commit(); err != nil {
		responses.SendError(context, http.StatusInternalServerError, "Erro ao confirmar transação")
		return
	}

	responses.SendSuccess(context, "delete-event", nil)
}
