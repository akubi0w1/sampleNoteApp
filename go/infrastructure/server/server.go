package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Addr string
	Port string
}

// NewServer サーバの作成
func NewServer(addr, port string) *Server {
	return &Server{
		Addr: addr,
		Port: port,
	}
}

// Serve サーバー起動
func (s *Server) Serve() {
	log.Println("Server running...")
	http.ListenAndServe(fmt.Sprintf("%s:%s", s.Addr, s.Port), nil)
}

// Get getリクエストを処理
func (s *Server) Get(endpoint string, apiFunc http.HandlerFunc) {
	http.HandleFunc(endpoint, httpMethod(apiFunc, http.MethodGet))
}

// Post postリクエストを処理
func (s *Server) Post(endpoint string, apiFunc http.HandlerFunc) {
	http.HandleFunc(endpoint, httpMethod(apiFunc, http.MethodPost))
}

func httpMethod(apiFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// CORS対応
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")

		// プリフライトリクエストは処理を通さない
		if request.Method == http.MethodOptions {
			return
		}

		// 指定のHTTPメソッドでない場合はエラー
		if request.Method != method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method not Allowed"))
		}

		// 共通のレスポンスヘッダを設定
		writer.Header().Add("Content-Type", "application/json")
		apiFunc(writer, request)
	}
}
