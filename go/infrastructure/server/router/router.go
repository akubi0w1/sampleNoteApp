package router

import (
	"note-app/infrastructure/server"
	"note-app/infrastructure/server/handler"
)

// SetUpRouting ルーティング
func SetUpRouting(s *server.Server, h *handler.AppHandler) {
	s.Get("/users/", h.UserHandler.GetUserByID)
}
