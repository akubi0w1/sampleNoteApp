package handler

import (
	"app/pkg/infrastructure/dcontext"
	"app/pkg/infrastructure/server/response"
	"app/pkg/interface/controller"
	"app/pkg/interface/repository"
	"app/pkg/usecase"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type noteHandler struct {
	NoteController controller.NoteController
}

type NoteHandler interface {
	GetNoteByNoteID(w http.ResponseWriter, r *http.Request)
	GetNotes(w http.ResponseWriter, r *http.Request)
	CreateNote(w http.ResponseWriter, r *http.Request)
	UpdateNote(w http.ResponseWriter, r *http.Request)
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
	noteID := strings.TrimPrefix(r.URL.Path, "/notes/")

	res, err := nh.NoteController.ShowNoteByNoteID(noteID)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	response.Success(w, res)
}

func (nh *noteHandler) GetNotes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := dcontext.GetUserIDFromContext(ctx)

	res, err := nh.NoteController.ShowNotes(userID)
	if err != nil {
		response.InternalServerError(w, err.Error())
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
		response.BadRequest(w, err.Error())
		return
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	res, err := nh.NoteController.CreateNote(userID, req)
	if err != nil {
		response.InternalServerError(w, err.Error())
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
		response.BadRequest(w, err.Error())
		return
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	res, err := nh.NoteController.UpdateNote(userID, noteID, req)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}
	response.Success(w, res)
}
