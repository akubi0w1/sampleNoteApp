package usecase

import (
	"app/pkg/domain"
)

type accountInteractor struct {
	AccountRepository AccountRepository
}

type AccountInteractor interface {
	Account(userID string) (domain.User, error)
	Auth(userID string) (domain.User, error)
}

func NewAccountInteractor(ar AccountRepository) AccountInteractor {
	return &accountInteractor{
		AccountRepository: ar,
	}
}

func (ai *accountInteractor) Account(userID string) (domain.User, error) {
	return ai.AccountRepository.FindByID(userID)
}

func (ai *accountInteractor) Auth(userID string) (domain.User, error) {
	return ai.AccountRepository.FindAuthByID(userID)
}
