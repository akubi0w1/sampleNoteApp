package handler

import (
	"app/pkg/infrastructure/server/response"
	"app/pkg/interface/controller"
	"app/pkg/interface/repository"
	"app/pkg/usecase"
	"net/http"
	"strings"
)

type userHandler struct {
	UserController controller.UserController
}

type UserHandler interface {
	GetUserByUserID(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(sh repository.SQLHandler) UserHandler {
	return &userHandler{
		UserController: controller.NewUserController(
			usecase.NewUserInteractor(
				repository.NewUserRepository(sh),
			),
		),
	}
}

func (uh *userHandler) GetUserByUserID(w http.ResponseWriter, r *http.Request) {
	userID := strings.TrimPrefix(r.URL.Path, "/users/")
	if userID == "" {
		response.BadRequest(w, "userID is empty")
		return
	}

	res, err := uh.UserController.ShowUserByUserID(userID)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	response.Success(w, res)
}

func (uh *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	res, err := uh.UserController.ShowUsers()
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}
	response.Success(w, res)
}
