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

func CreateSectorHandler(context *gin.Context) {

	db := database.GetDB()
	request := dto.SectorCreateDTO{}

	if err := context.ShouldBindJSON(&request); err != nil {

		responses.SendError(context, http.StatusBadRequest, err.Error())
		return
	}
	if err := request.Validade(); err != nil {
		responses.SendError(context, http.StatusBadRequest, err.Error())
		return
	}
	Sector := dto.SectorCreateDTO{
		Name:     request.Name,
		Capacity: request.Capacity,
		ShowID:   request.ShowID,
	}
	if err := repository.CreateSector(db, Sector); err != nil {
		responses.SendError(context, http.StatusInternalServerError, "nao foi possivel criar usuario")
		return
	}
	responses.SendSuccess(context, "create-sector", Sector)

}

func DeleteSectorHandler(context *gin.Context) {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		responses.SendError(context, http.StatusBadRequest, "id invalido")
		return
	}
	db := database.GetDB()
	if err := repository.DeleteSector(db, id); err != nil {
		responses.SendError(context, http.StatusInternalServerError, "nao foi possivel deletar usuario")
		return
	}
	responses.SendSuccess(context, "delete-user", nil)
}

func GetSectorByEventIdHandler(context *gin.Context) {
	idParam := context.Param("id")
	idEvent, err := strconv.Atoi(idParam)
	if err != nil {
		responses.SendError(context, http.StatusBadRequest, "id invalido")
		return
	}
	db := database.GetDB()

	sectors, err := repository.GetSectorsByEventID(db, idEvent)

	if err != nil {
		responses.SendError(context, http.StatusBadRequest, "erro ao buscar setores do evento")
		return
	}
	responses.SendSuccess(context, "get-sectors-by-event-id", sectors)

}
func UpdateTicketSectorHandler(context *gin.Context) {

	request := dto.UpdateSectorDTO{}
	if err := context.ShouldBindJSON(&request); err != nil {

		responses.SendError(context, http.StatusBadRequest, err.Error())
		return
	}
	if err := request.Validade(); err != nil {
		responses.SendError(context, http.StatusBadRequest, err.Error())
		return
	}
	UpdateSector := dto.UpdateSectorDTO{
		TicketId:    request.TicketId,
		NewSectorId: request.NewSectorId,
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

	ticket, err := repository.UpdateTicketSector(transaction, UpdateSector.TicketId, UpdateSector.NewSectorId)
	if err != nil {
		responses.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	if err := transaction.Commit(); err != nil {
		responses.SendError(context, http.StatusInternalServerError, "Erro ao confirmar transação")
		return
	}

	responses.SendSuccess(context, "Setor atualizado com sucesso", ticket)
}
