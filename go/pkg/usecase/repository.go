package usecase

import (
	"app/pkg/domain"
)

type AccountRepository interface {
	FindByID(userID string) (domain.User, error)
	FindAuthByID(userID string) (domain.User, error)
}
