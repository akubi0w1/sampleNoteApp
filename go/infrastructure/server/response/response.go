package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// Success
func Success(w http.ResponseWriter, response interface{}) {
	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		InternalServerError(w, "json marshal error")
		return
	}
	w.Write(jsonData)
}

// BadRequest
func BadRequest(w http.ResponseWriter, message string) {
	httpError(w, http.StatusBadRequest, message)
}

// InternalServerError
func InternalServerError(w http.ResponseWriter, message string) {
	httpError(w, http.StatusInternalServerError, message)
}

func httpError(w http.ResponseWriter, code int, message string) {
	jsonData, _ := json.Marshal(errorResponse{
		Code:    code,
		Message: message,
	})
	w.WriteHeader(code)
	if jsonData != nil {
		w.Write(jsonData)
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
