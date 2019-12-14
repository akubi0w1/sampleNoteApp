package controller

import (
	"note-app/usecase"
)

type UserController struct {
	UserInteractor usecase.UserInteractor
}

func (uc *UserController) Users() (*GetUsersResponse, error) {
	users, err := uc.UserInteractor.ShowUsers()
	if err != nil {
		return &GetUsersResponse{}, err
	}
	var response GetUsersResponse
	for _, v := range users {
		var res GetUserResponse
		res.ID = v.ID
		res.Name = v.Name
		res.Mail = v.Mail
		res.CreatedAt = v.CreatedAt
		response.Users = append(response.Users, res)
	}
	return &response, nil
}

func (uc *UserController) UserByID(userID string) (*GetUserResponse, error) {
	user, err := uc.UserInteractor.ShowUserByID(userID)
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

type GetUsersResponse struct {
	Users []GetUserResponse `json:"users"`
}

type GetUserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Mail      string `json:"mail"`
	CreatedAt string `json:"created_at"`
}
