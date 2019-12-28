package usecase

import (
	"app/pkg/domain"
	"errors"
	"github.com/google/uuid"
	"time"
)

type noteInteractor struct {
	NoteRepository NoteRepository
}

type NoteInteractor interface {
	NoteByNoteID(userID, noteID string) (domain.Note, error)
	Notes(userID string) (domain.Notes, error)
	AddNote(title, content, userID string) (note domain.Note, err error)
	UpdateNote(userID, id, title, content string) (note domain.Note, err error)
	DeleteNote(userID, noteID string) error
}

func NewNoteInteractor(nr NoteRepository) NoteInteractor {
	return &noteInteractor{
		NoteRepository: nr,
	}
}

func (ni *noteInteractor) NoteByNoteID(userID, noteID string) (note domain.Note, err error) {
	note, err = ni.NoteRepository.FindNoteByNoteID(noteID)
	if err != nil {
		return
	}
	if note.Author.ID != userID {
		return note, domain.Unauthorized(errors.New("Unauthorized"))
	}
	return
}

func (ni *noteInteractor) Notes(userID string) (domain.Notes, error) {
	return ni.NoteRepository.FindNotes(userID)
}

func (ni *noteInteractor) AddNote(title, content, userID string) (note domain.Note, err error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return note, domain.InternalServerError(err)
	}
	createdAt := time.Now()

	err = ni.NoteRepository.StoreNote(id.String(), title, content, userID, createdAt)
	if err != nil {
		return
	}
	return ni.NoteRepository.FindNoteByNoteID(id.String())
}

func (ni *noteInteractor) UpdateNote(userID, id, title, content string) (note domain.Note, err error) {
	updatedAt := time.Now()
	note, err = ni.NoteRepository.FindNoteByNoteID(id)
	if note.Author.ID != userID {
		return note, domain.Unauthorized(errors.New("Unauthorized"))
	}

	err = ni.NoteRepository.UpdateNote(id, title, content, updatedAt)
	if err != nil {
		return
	}
	note.Title = title
	note.Content = content
	note.UpdatedAt = updatedAt.String()
	return
}

func (ni *noteInteractor) DeleteNote(userID, noteID string) error {
	note, err := ni.NoteRepository.FindNoteByNoteID(noteID)
	if err != nil {
		return err
	}
	if note.Author.ID != userID {
		return domain.Unauthorized(errors.New("Unauthorized"))
	}
	return ni.NoteRepository.DeleteNote(noteID)
}
