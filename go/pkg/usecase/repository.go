package usecase

import (
	"time"

	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/domain"
)

type AccountRepository interface {
	FindByID(userID string) (domain.User, error)
	FindAuthByID(userID string) (domain.User, error)
	Store(userID, name, passoword, mail string, createdAt time.Time) error
	Update(userID, newUserID, name, password, mail string) error
	Delete(userID string) error
}

type UserRepository interface {
	FindUserByUserID(userID string) (domain.User, error)
	FindUsers() (domain.Users, error)
}

type NoteRepository interface {
	FindNoteByNoteID(noteID string) (domain.Note, error)
	FindNotes(userID string) (domain.Notes, error)
	StoreNote(id, title, content, userID string, createdAt time.Time) error
	UpdateNote(id, title, content string, updatedAt time.Time) error
	DeleteNote(noteID string) error
}
