package router

import (
	"app/pkg/infrastructure/server"
	"app/pkg/infrastructure/server/handler"
	"app/pkg/infrastructure/server/middleware"
	"net/http"
)

type router struct {
}

type Router interface {
}

func SetupRouter(serv server.Server, h handler.AppHandler) {
	http.HandleFunc("/accounts", middleware.Authorized(h.GetAccount()))
	http.HandleFunc("/login", h.Login())
}
