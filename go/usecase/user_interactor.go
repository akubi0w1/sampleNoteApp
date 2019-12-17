package usecase

import (
	"note-app/domain"
	"time"
	// "github.com/google/uuid"
)

type userInteractor struct {
	UserRepository UserRepository
}

type UserInteractor interface {
	Add(string, string, string, string) (domain.User, error)
	Delete(string) error
	ShowUsers() (domain.Users, error)
	ShowUserByID(string) (domain.User, error)
}

func NewUserInteractor(ur UserRepository) UserInteractor {
	return &userInteractor{
		UserRepository: ur,
	}
}

func (ui *userInteractor) Add(id, name, password, mail string) (user domain.User, err error) {
	user.ID = id
	user.Name = name
	user.Mail = mail
	user.Password = password
	createdAt := time.Now()
	user.CreatedAt = createdAt.String()

	err = ui.UserRepository.Store(id, name, password, mail, createdAt)
	return
}

func (ui *userInteractor) Delete(userID string) error {
	return ui.UserRepository.Delete(userID)
}

func (ui *userInteractor) ShowUsers() (domain.Users, error) {
	return ui.UserRepository.FindUsers()
}

func (ui *userInteractor) ShowUserByID(userID string) (domain.User, error) {
	return ui.UserRepository.FindUserByID(userID)
}
