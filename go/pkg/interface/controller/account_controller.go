package controller

import (
	"app/pkg/usecase"
	"errors"
)

type accountController struct {
	AccountInteractor usecase.AccountInteractor
}

type AccountController interface {
	ShowAccount(userID string) (*getAccountResponse, error)
	AuthAccount(userID string) (*loginResponse, error)
	Add(userID, name, password, mail string) (*createAccountResponse, error)
	UpdateAccount(userID string, req CreateAccountRequest) (*UpdateAccountResponse, error)
	DeleteAccount(userID string) error
}

func NewAccountController(ai usecase.AccountInteractor) AccountController {
	return &accountController{
		AccountInteractor: ai,
	}
}

func (ac *accountController) ShowAccount(userID string) (*getAccountResponse, error) {
	user, err := ac.AccountInteractor.Account(userID)
	if err != nil {
		return &getAccountResponse{}, err
	}
	return &getAccountResponse{
		UserID:    user.ID,
		Name:      user.Name,
		Mail:      user.Mail,
		CreatedAt: user.CreatedAt,
	}, nil
}

type getAccountResponse struct {
	UserID    string `json:"id"`
	Name      string `json:"name"`
	Mail      string `json:"mail"`
	CreatedAt string `json:"created_at"`
}

func (ac *accountController) Add(userID, name, password, mail string) (*createAccountResponse, error) {
	var res createAccountResponse
	if userID == "" {
		return &res, errors.New("userID is empty")
	}
	if name == "" {
		return &res, errors.New("name is empty")
	}
	if password == "" {
		return &res, errors.New("password is empty")
	}
	if mail == "" {
		return &res, errors.New("mail is empty")
	}

	user, err := ac.AccountInteractor.CreateAccount(userID, name, password, mail)
	if err != nil {
		return &res, err
	}
	res.ID = user.ID
	res.Name = user.Name
	res.Mail = user.Mail
	res.CreatedAt = user.CreatedAt
	return &res, nil
}

type CreateAccountRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type createAccountResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Mail      string `json:"mail"`
	CreatedAt string `json:"created_at"`
}

func (ac *accountController) UpdateAccount(userID string, req CreateAccountRequest) (*UpdateAccountResponse, error) {
	var res UpdateAccountResponse
	user, err := ac.AccountInteractor.UpdateAccount(userID, req.ID, req.Name, req.Password, req.Mail)
	if err != nil {
		return &res, err
	}
	res.ID = user.ID
	res.Name = user.Name
	res.Mail = user.Mail
	res.CreatedAt = user.CreatedAt
	return &res, nil
}

type UpdateAccountRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type UpdateAccountResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Mail      string `json:"mail"`
	CreatedAt string `json:"created_at"`
}

func (ac *accountController) DeleteAccount(userID string) error {
	return ac.AccountInteractor.DeleteAccount(userID)
}

// AuthAccount userIDの認証?
func (ac *accountController) AuthAccount(userID string) (*loginResponse, error) {
	user, err := ac.AccountInteractor.Auth(userID)
	if err != nil {
		return &loginResponse{}, err
	}
	return &loginResponse{
		UserID:   user.ID,
		Password: user.Password,
	}, nil
}

type LoginRequest struct {
	UserID   string `json:"id"`
	Password string `json:"password"`
}

type loginResponse struct {
	UserID   string `json:"id"`
	Password string `json:"password"`
}
