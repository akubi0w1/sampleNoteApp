package repository

import (
	"time"

	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/domain"
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/usecase"
)

type noteRepository struct {
	DB SQLHandler
}

func NewNoteRepository(sh SQLHandler) usecase.NoteRepository {
	return &noteRepository{
		DB: sh,
	}
}

func (nr *noteRepository) FindNoteByNoteID(noteID string) (note domain.Note, err error) {
	row := nr.DB.QueryRow("SELECT notes.id, title, content, notes.created_at, updated_at, users.id, name, mail, users.created_at FROM notes INNER JOIN users ON notes.user_id = users.id WHERE notes.id=?", noteID)
	if err = row.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt, &note.Author.ID, &note.Author.Name, &note.Author.Mail, &note.Author.CreatedAt); err != nil {
		return note, domain.InternalServerError(err)
	}
	return
}

func (nr *noteRepository) FindNotes(userID string) (notes domain.Notes, err error) {
	rows, err := nr.DB.Query("SELECT notes.id, title, content, notes.created_at, updated_at, users.id, name, mail, users.created_at FROM notes INNER JOIN users ON notes.user_id = users.id WHERE user_id=?", userID)
	if err != nil {
		return notes, domain.InternalServerError(err)
	}
	for rows.Next() {
		var note domain.Note
		if err = rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt, &note.Author.ID, &note.Author.Name, &note.Author.Mail, &note.Author.CreatedAt); err != nil {
			continue
		}
		notes = append(notes, note)
	}
	return
}

func (nr *noteRepository) StoreNote(id, title, content, userID string, created_at time.Time) error {
	_, err := nr.DB.Execute(
		"INSERT INTO notes(id, title, content, created_at, updated_at, user_id) VALUES (?,?,?,?,?,?)",
		id,
		title,
		content,
		created_at,
		created_at,
		userID,
	)
	return domain.InternalServerError(err)
}

func (nr *noteRepository) UpdateNote(id, title, content string, updatedAt time.Time) error {
	query := "UPDATE notes SET"
	var values []interface{}
	if title != "" {
		query += " title=?,"
		values = append(values, title)
	}
	query += " content=?, updated_at=? WHERE id=?"
	values = append(values, content, updatedAt, id)
	_, err := nr.DB.Execute(query, values...)
	return domain.InternalServerError(err)
}

func (nr *noteRepository) DeleteNote(noteID string) error {
	_, err := nr.DB.Execute("DELETE FROM notes WHERE id=?", noteID)
	return domain.InternalServerError(err)
}
