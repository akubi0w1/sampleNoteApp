package main

import (
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/infrastructure/database"
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/infrastructure/server"
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/infrastructure/server/handler"
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/infrastructure/server/router"
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
