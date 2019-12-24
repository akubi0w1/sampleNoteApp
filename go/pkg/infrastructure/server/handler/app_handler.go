package handler

import (
	"app/pkg/interface/repository"
	"net/http"
)

type appHandler struct {
	AccountHandler AccountHandler
}

type AppHandler interface {
	GetAccount() http.HandlerFunc
	Login() http.HandlerFunc
}

func NewAppHandler(sh repository.SQLHandler) AppHandler {
	return &appHandler{
		AccountHandler: NewAccountHandler(sh),
	}
}

func (ah *appHandler) GetAccount() http.HandlerFunc {
	return ah.AccountHandler.AccountHandler
}

func (ah *appHandler) Login() http.HandlerFunc {
	return ah.AccountHandler.Login
}
