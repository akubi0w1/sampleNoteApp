package handler

import (
	"note-app/interface/datastore"
)

// AppHandler アプリケーションハンドラ
type AppHandler struct {
	UserHandler UserHandler
}

// NewAppHandler アプリケーションハンドラの作成
func NewAppHandler(sh datastore.SQLHandler) *AppHandler {
	return &AppHandler{
		UserHandler: NewUserHandler(sh),
	}
}
