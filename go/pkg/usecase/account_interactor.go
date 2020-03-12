package usecase

import (
	"time"

	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/domain"
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/infrastructure/auth"
)

type accountInteractor struct {
	AccountRepository AccountRepository
}

type AccountInteractor interface {
	Account(userID string) (domain.User, error)
	Auth(userID string) (domain.User, error)
	CreateAccount(userID, name, password, mail string) (user domain.User, err error)
	UpdateAccount(userID, newUserID, name, password, mail string) (domain.User, error)
	DeleteAccount(userID string) error
}

func NewAccountInteractor(ar AccountRepository) AccountInteractor {
	return &accountInteractor{
		AccountRepository: ar,
	}
}

func (ai *accountInteractor) Account(userID string) (domain.User, error) {
	return ai.AccountRepository.FindByID(userID)
}

func (ai *accountInteractor) CreateAccount(userID, name, password, mail string) (user domain.User, err error) {
	// パスワードの暗号
	// TODO: 依存取って
	hash, err := auth.PasswordHash(password)
	if err != nil {
		return user, domain.InternalServerError(err)
	}

	// timeの取得
	createdAt := time.Now()

	err = ai.AccountRepository.Store(userID, name, hash, mail, createdAt)
	if err != nil {
		return
	}
	user.ID = userID
	user.Name = name
	user.Mail = mail
	user.CreatedAt = createdAt.String()
	return
}

func (ai *accountInteractor) UpdateAccount(userID, newUserID, name, password, mail string) (domain.User, error) {
	// TODO: 依存
	var hash string
	var err error
	var user domain.User
	if password != "" {
		hash, err = auth.PasswordHash(password)
		if err != nil {
			return user, domain.InternalServerError(err)
		}
	}

	err = ai.AccountRepository.Update(userID, newUserID, name, hash, mail)
	if err != nil {
		return user, err
	}

	user.ID = userID
	if newUserID != "" {
		user.ID = newUserID
	}
	user, err = ai.AccountRepository.FindByID(userID)
	return user, err
}

func (ai *accountInteractor) DeleteAccount(userID string) error {
	return ai.AccountRepository.Delete(userID)
}

func (ai *accountInteractor) Auth(userID string) (domain.User, error) {
	return ai.AccountRepository.FindAuthByID(userID)
}
