package repository

import (
	"app/pkg/domain"
	"app/pkg/usecase"
	"strings"
	"time"
)

type accountRepository struct {
	SQLHandler SQLHandler
}

func NewAccountRepository(sh SQLHandler) usecase.AccountRepository {
	return &accountRepository{
		SQLHandler: sh,
	}
}

func (ar *accountRepository) FindByID(userID string) (user domain.User, err error) {
	row := ar.SQLHandler.QueryRow("SELECT id, name, mail, created_at FROM users WHERE id=?", userID)
	if err = row.Scan(&user.ID, &user.Name, &user.Mail, &user.CreatedAt); err != nil {
		return user, domain.InternalServerError(err)
	}
	return
}

func (ar *accountRepository) FindAuthByID(userID string) (user domain.User, err error) {
	row := ar.SQLHandler.QueryRow("SELECT id, password FROM users WHERE id=?", userID)
	if err = row.Scan(&user.ID, &user.Password); err != nil {
		return user, domain.InternalServerError(err)
	}
	return
}

func (ar *accountRepository) Store(userID, name, passoword, mail string, createdAt time.Time) error {
	_, err := ar.SQLHandler.Execute(
		"INSERT INTO users(id, name, password, mail, created_at) VALUES (?,?,?,?,?)",
		userID,
		name,
		passoword,
		mail,
		createdAt,
	)
	return domain.InternalServerError(err)
}

func (ar *accountRepository) Update(userID, newUserID, name, password, mail string) error {
	query := "UPDATE users SET"
	var values []interface{}
	if newUserID != "" {
		query += " id=?,"
		values = append(values, newUserID)
	}
	if name != "" {
		query += " name=?,"
		values = append(values, name)
	}
	if mail != "" {
		query += " mail=?,"
		values = append(values, mail)
	}
	if password != "" {
		query += " password=?,"
		values = append(values, password)
	}
	query = strings.TrimSuffix(query, ",")
	query += " WHERE id=?;"
	values = append(values, userID)
	_, err := ar.SQLHandler.Execute(query, values...)
	return domain.InternalServerError(err)
}

func (ar *accountRepository) Delete(userID string) error {
	_, err := ar.SQLHandler.Execute("DELETE FROM users WHERE id=?", userID)
	return domain.InternalServerError(err)
}
