package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"note-app/infrastructure/server/response"
	"note-app/interface/controller"
	"note-app/interface/datastore"
	"note-app/usecase"
)

// UserHandler
type userHandler struct {
	UserController controller.UserController
}

type UserHandler interface {
	CreateUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
	GetUsers(http.ResponseWriter, *http.Request)
	GetUserByID(http.ResponseWriter, *http.Request)
}

// NewUserHandler
func NewUserHandler(sh datastore.SQLHandler) UserHandler {
	return &userHandler{
		UserController: controller.NewUserController(
			usecase.NewUserInteractor(
				datastore.NewUserRepository(
					sh,
				),
			),
		),
	}
}

func (uh *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	var req controller.CreateAccountRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	res, err := uh.UserController.Create(&req)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}
	response.Success(w, res)
}

func (uh *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// とりあえずpathで検出...
	userID := strings.TrimPrefix(r.URL.Path, "/account/delete/")

	res, err := uh.UserController.Delete(userID)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}
	response.Success(w, res)
}

func (uh *userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	res, err := uh.UserController.Users()
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}
	response.Success(w, res)
}

func (uh *userHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := strings.TrimPrefix(r.URL.Path, "/users/")

	res, err := uh.UserController.UserByID(userID)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	response.Success(w, res)
}
