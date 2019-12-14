package datastore

import (
	"note-app/domain"
)

type UserRepository struct {
	SQLHandler SQLHandler
}

func (ur *UserRepository) FindUsers() (users domain.Users, err error) {
	rows, err := ur.SQLHandler.Query("SELECT id, name, mail, created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		var id string
		var name string
		var mail string
		var createdAt string
		if err = rows.Scan(&id, &name, &mail, &createdAt); err != nil {
			continue
		}
		users = append(users, domain.User{ID: id, Name: name, Mail: mail, CreatedAt: createdAt})
	}
	return

}

func (ur *UserRepository) FindUserByID(userID string) (user domain.User, err error) {
	row := ur.SQLHandler.QueryRow("SELECT id, name, mail, created_at FROM users WHERE id=?", userID)
	if err = row.Scan(&user.ID, &user.Name, &user.Mail, &user.CreatedAt); err != nil {
		return
	}
	return
}
