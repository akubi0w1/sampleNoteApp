package controller

import (
	"errors"

	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/domain"
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/usecase"
)

type noteController struct {
	NoteInteractor usecase.NoteInteractor
}

type NoteController interface {
	ShowNoteByNoteID(userID, noteID string) (*GetNoteResponse, error)
	ShowNotes(userID string) (*GetNotesResopnse, error)
	CreateNote(userID string, req CreateNoteRequest) (*CreateNoteResponse, error)
	UpdateNote(userID, noteID string, req UpdateNoteRequest) (*UpdateNoteResponse, error)
	DeleteNote(userID, noteID string) error
}

func NewNoteController(ni usecase.NoteInteractor) NoteController {
	return &noteController{
		NoteInteractor: ni,
	}
}

func (nc *noteController) ShowNoteByNoteID(userID, noteID string) (*GetNoteResponse, error) {
	var res GetNoteResponse
	note, err := nc.NoteInteractor.NoteByNoteID(userID, noteID)
	if err != nil {
		return &res, err
	}
	res.ID = note.ID
	res.Title = note.Title
	res.Content = note.Content
	res.CreatedAt = note.CreatedAt
	res.UpdatedAt = note.UpdatedAt
	res.Author.ID = note.Author.ID
	res.Author.Name = note.Author.Name
	res.Author.Mail = note.Author.Mail
	res.Author.CreatedAt = note.Author.CreatedAt
	return &res, err
}

type GetNoteResponse struct {
	ID        string          `json:"id"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
	Author    GetUserResponse `json:"author"`
}

func (nc *noteController) ShowNotes(userID string) (*GetNotesResopnse, error) {
	var res GetNotesResopnse
	notes, err := nc.NoteInteractor.Notes(userID)
	if err != nil {
		return &res, err
	}
	for _, note := range notes {
		res.Notes = append(res.Notes, GetNoteResponse{
			ID:        note.ID,
			Title:     note.Title,
			Content:   note.Content,
			CreatedAt: note.CreatedAt,
			UpdatedAt: note.UpdatedAt,
			Author: GetUserResponse{
				ID:        note.Author.ID,
				Name:      note.Author.Name,
				Mail:      note.Author.Mail,
				CreatedAt: note.Author.CreatedAt,
			},
		})
	}
	return &res, nil
}

type GetNotesResopnse struct {
	Notes []GetNoteResponse `json:"notes"`
}

func (nc *noteController) CreateNote(userID string, req CreateNoteRequest) (*CreateNoteResponse, error) {
	var res CreateNoteResponse
	if req.Title == "" {
		return &res, domain.BadRequest(errors.New("title is empty"))
	}
	if req.Content == "" {
		return &res, domain.BadRequest(errors.New("content is empty"))
	}
	note, err := nc.NoteInteractor.AddNote(req.Title, req.Content, userID)
	if err != nil {
		return &res, err
	}
	res.ID = note.ID
	res.Title = note.Title
	res.Content = note.Content
	res.CreatedAt = note.CreatedAt
	res.UpdatedAt = note.UpdatedAt
	res.Author.ID = note.Author.ID
	res.Author.Name = note.Author.Name
	res.Author.Mail = note.Author.Mail
	res.Author.CreatedAt = note.Author.CreatedAt
	return &res, nil
}

type CreateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateNoteResponse struct {
	ID        string          `json:"id"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
	Author    GetUserResponse `json:"author"`
}

func (nc *noteController) UpdateNote(userID, noteID string, req UpdateNoteRequest) (*UpdateNoteResponse, error) {
	var res UpdateNoteResponse
	if req.Title == "" {
		return &res, domain.BadRequest(errors.New("title is empty"))
	}
	note, err := nc.NoteInteractor.UpdateNote(userID, noteID, req.Title, req.Content)
	if err != nil {
		return &res, err
	}
	res.ID = note.ID
	res.Title = note.Title
	res.Content = note.Content
	res.CreatedAt = note.CreatedAt
	res.UpdatedAt = note.UpdatedAt
	res.Author.ID = note.Author.ID
	res.Author.Name = note.Author.Name
	res.Author.Mail = note.Author.Mail
	res.Author.CreatedAt = note.Author.CreatedAt
	return &res, nil
}

type UpdateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateNoteResponse struct {
	ID        string          `json:"id"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	CreatedAt string          `json:"created_at"`
	UpdatedAt string          `json:"updated_at"`
	Author    GetUserResponse `json:"author"`
}

func (nc *noteController) DeleteNote(userID, noteID string) error {
	return nc.NoteInteractor.DeleteNote(userID, noteID)
}
