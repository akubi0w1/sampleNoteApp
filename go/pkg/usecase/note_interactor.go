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
	NoteByNoteID(noteID string) (domain.Note, error)
	Notes(userID string) (domain.Notes, error)
	AddNote(title, content, userID string) (note domain.Note, err error)
	UpdateNote(userID, id, title, content string) (note domain.Note, err error)
}

func NewNoteInteractor(nr NoteRepository) NoteInteractor {
	return &noteInteractor{
		NoteRepository: nr,
	}
}

func (ni *noteInteractor) NoteByNoteID(noteID string) (domain.Note, error) {
	return ni.NoteRepository.FindNoteByNoteID(noteID)
}

func (ni *noteInteractor) Notes(userID string) (domain.Notes, error) {
	return ni.NoteRepository.FindNotes(userID)
}

func (ni *noteInteractor) AddNote(title, content, userID string) (note domain.Note, err error) {
	// gen note id
	id, err := uuid.NewRandom()
	if err != nil {
		return
	}
	// get time
	createdAt := time.Now()
	// store db
	err = ni.NoteRepository.StoreNote(id.String(), title, content, userID, createdAt)
	if err != nil {
		return
	}
	return ni.NoteRepository.FindNoteByNoteID(id.String())
}

func (ni *noteInteractor) UpdateNote(userID, id, title, content string) (note domain.Note, err error) {
	// get time
	updatedAt := time.Now()
	note, err = ni.NoteRepository.FindNoteByNoteID(id)
	if note.Author.ID != userID {
		return note, errors.New("auth error")
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
