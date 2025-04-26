package handler

import (
	"ingressos-api/database"
	"ingressos-api/dto"
	"ingressos-api/repository"
	responses "ingressos-api/responses"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(context *gin.Context) {
	request := dto.CreateUserDTO{}
	db := database.GetDB()

	if err := context.ShouldBindJSON(&request); err != nil {

		responses.SendError(context, http.StatusBadRequest, "Dados inválidos.")
		return
	}

	if err := request.Validate(); err != nil {

		responses.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	user := dto.CreateUserDTO{
		Name:  request.Name,
		Email: request.Email,
	}

	if err := repository.CreateUserRepository(db, &user); err != nil {

		if strings.Contains(err.Error(), "duplicate key") {
			responses.SendError(context, http.StatusConflict, "Email já cadastrado.")
		} else {
			responses.SendError(context, http.StatusInternalServerError, "Erro ao criar usuário.")
		}
		return
	}

	responses.SendSuccess(context, "create-user", user)
}

func GetAllUsersHandler(context *gin.Context) {
	bd := database.GetDB()

	users, error := repository.GetAllUsersRepository(bd)
	if error != nil {
		responses.SendError(context, http.StatusInternalServerError, "erro ao buscar usuarios")
		return
	}
	responses.SendSuccess(context, "Get all users", users)

}

func GetUserByIdHandler(context *gin.Context) {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		responses.SendError(context, http.StatusBadRequest, "id invalido")
		return
	}
	bd := database.GetDB()

	user, error := repository.GetUserByID(bd, id)
	if error != nil {
		responses.SendError(context, http.StatusBadRequest, "erro ao buscar usuario")
		return
	}
	responses.SendSuccess(context, "get user by id", user)
}

func DeleteUserById(context *gin.Context) {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		responses.SendError(context, http.StatusBadRequest, "id invalido")
		return
	}
	db := database.GetDB()
	if err := repository.DeleteUserById(db, id); err != nil {
		responses.SendError(context, http.StatusBadRequest, "nao foi possivel deletar usuario")
		return
	}
	responses.SendSuccess(context, "delete-user", nil)
}
