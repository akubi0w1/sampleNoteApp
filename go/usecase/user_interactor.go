package usecase

import (
	"note-app/domain"
	"github.com/google/uuid"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (ui *UserInteractor) ShowUsers() (domain.Users, error) {
	return ui.UserRepository.FindUsers()
}

func (ui *UserInteractor) ShowUserByID(userID string) (domain.User, error) {
	return ui.UserRepository.FindUserByID(userID)
}
