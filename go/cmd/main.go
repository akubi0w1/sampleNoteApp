package main

import (
	"note-app/config"
	"note-app/infrastructure/database"
	"note-app/infrastructure/server"
	"note-app/infrastructure/server/handler"
	"note-app/infrastructure/server/router"
)

func main() {
	// load config
	servConf := config.LoadServerConfig()

	// connection db
	dbHandler := database.NewSQLHandler()

	// create handler
	appHandler := handler.NewAppHandler(dbHandler)

	// create new server
	serv := server.NewServer(servConf.Addr, servConf.Port)

	// routing
	router.SetUpRouting(serv, appHandler)

	// start app
	serv.Serve()
}
