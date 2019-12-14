package usecase

import (
	"note-app/domain"
)

type UserRepository interface {
	FindUsers() (domain.Users, error)
	FindUserByID(string) (domain.User, error)
}
