package usecase

import (
	"note-app/domain"
)

type UserRepository interface {
	FindUserByID(string) (domain.User, error)
}
