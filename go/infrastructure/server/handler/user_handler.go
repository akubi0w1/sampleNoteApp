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

// GetUserByID get user infomation by user id
func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := strings.TrimPrefix(r.URL.Path, "/users/")

	// TODO: dbにアクセス
	res, err := uh.UserController.UserByID(userID)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	// TODO: responseの作成
	response.Success(w, res)
}

// Hello sample handler
func Hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("heee"))
}
