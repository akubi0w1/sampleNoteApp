package datastore

import (
	"note-app/domain"
	"note-app/usecase"
	"time"
)

type userRepository struct {
	SQLHandler SQLHandler
}

func NewUserRepository(sh SQLHandler) usecase.UserRepository {
	return &userRepository{
		SQLHandler: sh,
	}
}

func (ur *userRepository) Store(id, name, password, mail string, createdAt time.Time) error {
	_, err := ur.SQLHandler.Execute(
		"INSERT INTO users(id, name, password, mail, created_at) VALUES (?,?,?,?,?)",
		id,
		name,
		password,
		mail,
		createdAt,
	)
	return err
}

func (ur *userRepository) Delete(userID string) error {
	_, err := ur.SQLHandler.Execute("DELETE FROM users WHERE id=?", userID)
	return err
}

func (ur *userRepository) FindUsers() (users domain.Users, err error) {
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

func (ur *userRepository) FindUserByID(userID string) (user domain.User, err error) {
	row := ur.SQLHandler.QueryRow("SELECT id, name, mail, created_at FROM users WHERE id=?", userID)
	if err = row.Scan(&user.ID, &user.Name, &user.Mail, &user.CreatedAt); err != nil {
		return
	}
	return
}
