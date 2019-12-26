package usecase

import (
	"app/pkg/domain"
)

type noteInteractor struct {
	NoteRepository NoteRepository
}

type NoteInteractor interface {
	NoteByNoteID(noteID string) (domain.Note, error)
	Notes(userID string) (domain.Notes, error)
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
