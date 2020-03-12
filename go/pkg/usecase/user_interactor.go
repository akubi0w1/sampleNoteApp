package usecase

import (
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/domain"
)

type userInteractor struct {
	UserRepository UserRepository
}

type UserInteractor interface {
	UserByUserID(userID string) (domain.User, error)
	Users() (domain.Users, error)
}

func NewUserInteractor(ur UserRepository) UserInteractor {
	return &userInteractor{
		UserRepository: ur,
	}
}

func (ui *userInteractor) UserByUserID(userID string) (domain.User, error) {
	return ui.UserRepository.FindUserByUserID(userID)
}

func (ui *userInteractor) Users() (domain.Users, error) {
	return ui.UserRepository.FindUsers()
}
