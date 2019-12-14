package datastore

import (
	"note-app/domain"
)

type UserRepository struct {
	SQLHandler SQLHandler
}

func (ur *UserRepository) FindUserByID(userID string) (user domain.User, err error) {
	row := ur.SQLHandler.QueryRow("SELECT id, name, mail, created_at FROM users WHERE id=?", userID)
	if err = row.Scan(&user.ID, &user.Name, &user.Mail, &user.CreatedAt); err != nil {
		return
	}
	return
}
