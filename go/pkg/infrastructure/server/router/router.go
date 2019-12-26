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
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")

		// プリフライトリクエストは処理を通さない
		if r.Method == http.MethodOptions {
			return
		}

		// 共通のレスポンスヘッダを設定
		w.Header().Add("Content-Type", "application/json")

		apiFunc(w, r)
	})
}
