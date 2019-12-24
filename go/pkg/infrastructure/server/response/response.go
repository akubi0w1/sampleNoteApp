package response

import (
	"encoding/json"
	"net/http"
)

func Success(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		InternalServerError(w, err.Error())
		return
	}
	w.Write(jsonData)
}

func BadRequest(w http.ResponseWriter, message string) {
	httpError(w, http.StatusBadRequest, message)
}

func InternalServerError(w http.ResponseWriter, message string) {
	httpError(w, http.StatusInternalServerError, message)
}

func httpError(w http.ResponseWriter, code int, message string) {
	jsonData, _ := json.Marshal(&errorResponse{
		Code:    code,
		Message: message,
	})
	w.Write(jsonData)
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
