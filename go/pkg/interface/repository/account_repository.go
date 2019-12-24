package repository

import (
	"app/pkg/domain"
	"app/pkg/usecase"
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
	row := ar.SQLHandler.QueryRow("SELECT id, name, mail FROM users WHERE id=?", userID)
	if err = row.Scan(&user.ID, &user.Name, &user.Mail); err != nil {
		return
	}
	return
}

func (ar *accountRepository) FindAuthByID(userID string) (user domain.User, err error) {
	row := ar.SQLHandler.QueryRow("SELECT id, password FROM users WHERE id=?", userID)
	if err = row.Scan(&user.ID, &user.Password); err != nil {
		return
	}
	return
}
