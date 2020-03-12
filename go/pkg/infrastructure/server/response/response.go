package response

import (
	"encoding/json"
	"net/http"

	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/domain"
)

func Success(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		domain.InternalServerError(err)
		return
	}
	w.Write(jsonData)
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func HttpError(w http.ResponseWriter, err error) {
	e, ok := err.(domain.Error)
	if !ok {
		e = domain.InternalServerError(err)
	}
	jsonData, _ := json.Marshal(&errorResponse{
		Code:    e.GetStatusCode(),
		Message: e.Error(),
	})
	w.WriteHeader(e.GetStatusCode())
	w.Write(jsonData)
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
