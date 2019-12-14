package usecase

import (
	"note-app/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (ui *UserInteractor) ShowUserByID(userID string) (domain.User, error) {
	return ui.UserRepository.FindUserByID(userID)
}
