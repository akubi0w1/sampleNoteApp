package controller

import (
	"app/pkg/usecase"
)

type userController struct {
	UserInteractor usecase.UserInteractor
}

type UserController interface {
	ShowUserByUserID(userID string) (*GetUserResponse, error)
	ShowUsers() (*getUsersResponse, error)
}

func NewUserController(ui usecase.UserInteractor) UserController {
	return &userController{
		UserInteractor: ui,
	}
}

func (uc *userController) ShowUserByUserID(userID string) (*GetUserResponse, error) {
	user, err := uc.UserInteractor.UserByUserID(userID)
	if err != nil {
		return &GetUserResponse{}, err
	}
	return &GetUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Mail:      user.Mail,
		CreatedAt: user.CreatedAt,
	}, nil
}

type GetUserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Mail      string `json:"mail"`
	CreatedAt string `json:"created_at"`
}

func (uc *userController) ShowUsers() (*getUsersResponse, error) {
	users, err := uc.UserInteractor.Users()
	var res getUsersResponse
	if err != nil {
		return &res, err
	}
	for _, user := range users {
		data := GetUserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Mail:      user.Mail,
			CreatedAt: user.CreatedAt,
		}
		res.Users = append(res.Users, data)
	}
	return &res, nil
}

type getUsersResponse struct {
	Users []GetUserResponse `json:"users"`
}
