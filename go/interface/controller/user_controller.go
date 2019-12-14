package controller

import (
	"note-app/usecase"
)

type UserController struct {
	UserInteractor usecase.UserInteractor
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

type GetUserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Mail      string `json:"mail"`
	CreatedAt string `json:"created_at"`
}
