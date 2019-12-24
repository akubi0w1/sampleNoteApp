package main

import (
	"app/pkg/infrastructure/database"
	"app/pkg/infrastructure/server"
	"app/pkg/infrastructure/server/handler"
	"app/pkg/infrastructure/server/router"
)

func main() {
	// TODO: load config

	// connect db
	sh := database.NewSQLHandler()

	// make server
	serv := server.NewServer("localhost", "8080")

	// make handler
	ah := handler.NewAppHandler(sh)

	// make router
	router.SetupRouter(serv, ah)

	// listen and serve
	serv.Serve()
}
