package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"app/pkg/domain"
	"app/pkg/infrastructure/auth"
	"app/pkg/infrastructure/dcontext"
	"app/pkg/infrastructure/server/logger"
	"app/pkg/infrastructure/server/response"
	"app/pkg/interface/controller"
	"app/pkg/interface/repository"
	"app/pkg/usecase"
)

type accountHandler struct {
	AccountController controller.AccountController
}

type AccountHandler interface {
	GetAccount(http.ResponseWriter, *http.Request)
	CreateAccount(w http.ResponseWriter, r *http.Request)
	UpdateAccount(w http.ResponseWriter, r *http.Request)
	DeleteAccount(w http.ResponseWriter, r *http.Request)
	Login(http.ResponseWriter, *http.Request)
	// LogoutHandler(http.ResponseWriter, *http.Request)
}

func NewAccountHandler(sh repository.SQLHandler) AccountHandler {
	return &accountHandler{
		AccountController: controller.NewAccountController(
			usecase.NewAccountInteractor(
				repository.NewAccountRepository(sh),
			),
		),
	}
}

func (ah *accountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	// contextからuserIDの取り出し
	ctx := r.Context()
	userID := dcontext.GetUserIDFromContext(ctx)
	if userID == "" {
		response.HttpError(w, domain.BadRequest(errors.New("userID is empty")))
		return
	}

	res, err := ah.AccountController.ShowAccount(userID)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, err)
		return
	}
	// レスポンスを作成
	response.Success(w, res)
}

func (ah *accountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// bodyの読み出し
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var req controller.CreateAccountRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	// controllerを叩く
	res, err := ah.AccountController.Add(req.ID, req.Name, req.Password, req.Mail)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, err)
		return
	}

	// response
	response.Success(w, res)
}

func (ah *accountHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	// ctxからuserIDの取得
	ctx := r.Context()
	userID := dcontext.GetUserIDFromContext(ctx)

	// bodyの読み出し
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var req controller.CreateAccountRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	res, err := ah.AccountController.UpdateAccount(userID, req)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, err)
		return
	}

	response.Success(w, res)
}

func (ah *accountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := dcontext.GetUserIDFromContext(ctx)

	err := ah.AccountController.DeleteAccount(userID)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, err)
		return
	}
	response.NoContent(w)
}

func (ah *accountHandler) Login(w http.ResponseWriter, r *http.Request) {
	// requestbodyを読む
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.HttpError(w, domain.BadRequest(err))
		return
	}
	var req controller.LoginRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	// TODO: usecaseに持っていく
	// userの取得
	res, err := ah.AccountController.AuthAccount(req.UserID)

	// パスワードの認証
	err = auth.PasswordVerify(res.Password, req.Password)
	if err != nil {
		logger.Warn(err)
		response.HttpError(w, domain.Unauthorized(err))
		return
	}

	// jwtの発行
	token, err := auth.CreateToken(req.UserID)
	if err != nil {
		logger.Error(err)
		response.HttpError(w, domain.InternalServerError(err))
		return
	}

	// クッキーに乗っける
	cookie := http.Cookie{
		Name:   "x-token",
		Value:  token,
		Domain: "localhost",
		Path:   "/",
	}
	http.SetCookie(w, &cookie)

	// TODO: resを受け取る

	// response
	response.Success(w, res)
}
