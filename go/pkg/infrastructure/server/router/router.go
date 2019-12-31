package router

import (
	"app/pkg/infrastructure/server"
	"app/pkg/infrastructure/server/handler"
	"net/http"
)

type router struct {
}

type Router interface {
}

func SetupRouter(serv server.Server, h handler.AppHandler) {
	Handle("/accounts", h.ManageAccount())
	Handle("/login", h.Login())

	Handle("/users", h.GetUsers())
	Handle("/users/", h.GetUserByUserID())

	Handle("/notes", h.ManageNotes())
	Handle("/notes/", h.ManageANote())
}

func Handle(endpoint string, apiFunc http.HandlerFunc) {
	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		// CORS対応
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000") // client server addr
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")

		// プリフライトリクエストは処理を通さない
		if r.Method == http.MethodOptions {
			return
		}

		// 共通のレスポンスヘッダを設定
		w.Header().Add("Content-Type", "application/json")

		apiFunc(w, r)
	})
}
