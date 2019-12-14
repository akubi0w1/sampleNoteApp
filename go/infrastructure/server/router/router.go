package router

import (
	"note-app/infrastructure/server"
	"note-app/infrastructure/server/handler"
)

func SetUpRouting(s *server.Server, h *handler.UserHandler) {
	// s.Get("/hello", handler.Hello)
	s.Get("/users/", h.GetUserByID)
}
