package main

import (
	"note-app/config"
	"note-app/infrastructure/database"
	"note-app/infrastructure/server"
	"note-app/infrastructure/server/handler"
	"note-app/infrastructure/server/router"
)

func main() {
	// TODO: load config
	servConf := config.LoadServerConfig()

	// TODO: connection db
	dbHandler := database.NewSQLHandler()

	// TODO: create handler
	userHandler := handler.NewUserHandler(dbHandler)

	// TODO: create new server
	serv := server.NewServer(servConf.Addr, servConf.Port)

	// TODO: routing
	router.SetUpRouting(serv, userHandler)

	// TODO: start app
	serv.Serve()
}
