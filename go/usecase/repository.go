package usecase

import (
	"note-app/domain"
	"time"
)

type UserRepository interface {
	Store(string, string, string, string, time.Time) error
	Delete(string) error
	FindUsers() (domain.Users, error)
	FindUserByID(string) (domain.User, error)
}
