package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/domain"
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/infrastructure/dcontext"
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/infrastructure/server/logger"
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/infrastructure/server/response"
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/interface/controller"
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/interface/repository"
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/usecase"
)

type noteHandler struct {
	NoteController controller.NoteController
}

type NoteHandler interface {
	GetNoteByNoteID(w http.ResponseWriter, r *http.Request)
	GetNotes(w http.ResponseWriter, r *http.Request)
	CreateNote(w http.ResponseWriter, r *http.Request)
	UpdateNote(w http.ResponseWriter, r *http.Request)
	DeleteNote(w http.ResponseWriter, r *http.Request)
}

func NewNoteHandler(sh repository.SQLHandler) NoteHandler {
	return &noteHandler{
		NoteController: controller.NewNoteController(
			usecase.NewNoteInteractor(
				repository.NewNoteRepository(sh),
			),
		),
	}
}

func (nh *noteHandler) GetNoteByNoteID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := dcontext.GetUserIDFromContext(ctx)

	noteID := strings.TrimPrefix(r.URL.Path, "/notes/")

	res, err := nh.NoteController.ShowNoteByNoteID(userID, noteID)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, err)
		return
	}

	response.Success(w, res)
}

func (nh *noteHandler) GetNotes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := dcontext.GetUserIDFromContext(ctx)

	res, err := nh.NoteController.ShowNotes(userID)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, err)
		return
	}
	response.Success(w, res)
}

func (nh *noteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := dcontext.GetUserIDFromContext(ctx)

	var req controller.CreateNoteRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	res, err := nh.NoteController.CreateNote(userID, req)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, err)
		return
	}
	response.Success(w, res)
}

func (nh *noteHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := dcontext.GetUserIDFromContext(ctx)

	noteID := strings.TrimPrefix(r.URL.Path, "/notes/")

	var req controller.UpdateNoteRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	res, err := nh.NoteController.UpdateNote(userID, noteID, req)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, err)
		return
	}
	response.Success(w, res)
}

func (nh *noteHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := dcontext.GetUserIDFromContext(ctx)

	noteID := strings.TrimPrefix(r.URL.Path, "/notes/")

	err := nh.NoteController.DeleteNote(userID, noteID)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, err)
		return
	}
	response.NoContent(w)
}
