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
type UserHandler struct {
	UserController controller.UserController
}

// NewUserHandler
func NewUserHandler(sh datastore.SQLHandler) *UserHandler {
	return &UserHandler{
		UserController: controller.UserController{
			UserInteractor: usecase.UserInteractor{
				UserRepository: &datastore.UserRepository{
					SQLHandler: sh,
				},
			},
		},
	}
}

func (uh *UserHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
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

func (uh *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// とりあえずpathで検出...
	userID := strings.TrimPrefix(r.URL.Path, "/account/delete/")

	res, err := uh.UserController.Delete(userID)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}
	response.Success(w, res)
}

func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	res, err := uh.UserController.Users()
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}
	response.Success(w, res)
}

func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := strings.TrimPrefix(r.URL.Path, "/users/")

	res, err := uh.UserController.UserByID(userID)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	response.Success(w, res)
}
