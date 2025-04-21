package handler

import (
	"ingressos-api/database"
	"ingressos-api/dto"
	"ingressos-api/repository"
	responses "ingressos-api/responses"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(context *gin.Context) {
	request := dto.CreateUserDTO{}
	db := database.GetDB()

	if err := context.ShouldBindJSON(&request); err != nil {
		logger.Errorf("bind error: %v", err.Error())
		responses.SendError(context, http.StatusBadRequest, "Dados inválidos.")
		return
	}

	// Validação do DTO
	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		responses.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	user := dto.CreateUserDTO{
		Name:  request.Name,
		Email: request.Email,
	}

	// Salvar no banco
	if err := repository.CreateUserRepository(db, &user); err != nil {
		logger.Errorf("error creating user: %v", err.Error())

		// Tratar erro de email duplicado
		if strings.Contains(err.Error(), "duplicate key") {
			responses.SendError(context, http.StatusConflict, "Email já cadastrado.")
		} else {
			responses.SendError(context, http.StatusInternalServerError, "Erro ao criar usuário.")
		}
		return
	}

	logger.Infof("usuário criado com sucesso: %v", user.Email)
	responses.SendSuccess(context, "create-user", user)
}

func GetAllUsersHandler() {

}

func GetUserByIdHandler() {

}
