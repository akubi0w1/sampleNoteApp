package usecase

import (
	"note-app/domain"
	"time"
	// "github.com/google/uuid"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (ui *UserInteractor) Add(id, name, password, mail string) (user domain.User, err error) {
	user.ID = id
	user.Name = name
	user.Mail = mail
	user.Password = password
	createdAt := time.Now()
	user.CreatedAt = createdAt.String()

	err = ui.UserRepository.Store(id, name, password, mail, createdAt)
	return
}

func (ui *UserInteractor) Remove(userID string) error {
	return ui.UserRepository.Delete(userID)
}

func (ui *UserInteractor) ShowUsers() (domain.Users, error) {
	return ui.UserRepository.FindUsers()
}

func (ui *UserInteractor) ShowUserByID(userID string) (domain.User, error) {
	return ui.UserRepository.FindUserByID(userID)
}
