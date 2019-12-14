package handler

import (
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

// TODO: define func: GetUsers
func (uh *UserHandler) GetUsers(w http.ResponseWriter, t *http.Request) {
	res, err := uh.UserController.Users()
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}
	response.Success(w, res)
}

// GetUserByID get user infomation by user id
func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := strings.TrimPrefix(r.URL.Path, "/users/")

	res, err := uh.UserController.UserByID(userID)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	response.Success(w, res)
}
