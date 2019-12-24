package server

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {
	Addr string
	Port string
}

type Server interface {
	Serve()
}

func NewServer(addr, port string) Server {
	return &server{
		Addr: addr,
		Port: port,
	}
}

func (s *server) Serve() {
	log.Println("Server running...")
	http.ListenAndServe(
		fmt.Sprintf("%s:%s", s.Addr, s.Port),
		nil,
	)
}
