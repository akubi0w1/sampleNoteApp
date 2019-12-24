package controller

import (
	"app/pkg/usecase"
)

type accountController struct {
	AccountInteractor usecase.AccountInteractor
}

type AccountController interface {
	ShowAccount(userID string) (*getAccountResponse, error)
	AuthAccount(userID string) (*loginResponse, error)
}

func NewAccountController(ai usecase.AccountInteractor) AccountController {
	return &accountController{
		AccountInteractor: ai,
	}
}

func (ac *accountController) ShowAccount(userID string) (*getAccountResponse, error) {
	user, err := ac.AccountInteractor.Account(userID)
	if err != nil {
		return &getAccountResponse{}, err
	}
	return &getAccountResponse{
		UserID: user.ID,
		Name:   user.Name,
		Mail:   user.Mail,
	}, nil
}

type getAccountResponse struct {
	UserID string `json:"id"`
	Name   string `json:"name"`
	Mail   string `json:"mail"`
}

func (ac *accountController) AuthAccount(userID string) (*loginResponse, error) {
	user, err := ac.AccountInteractor.Auth(userID)
	if err != nil {
		return &loginResponse{}, err
	}
	return &loginResponse{
		UserID:   user.ID,
		Password: user.Password,
	}, nil
}

type LoginRequest struct {
	UserID   string `json:"id"`
	Password string `json:"password"`
}

type loginResponse struct {
	UserID   string `json:"id"`
	Password string `json:"password"`
}
