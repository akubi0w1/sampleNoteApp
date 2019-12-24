package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"app/pkg/infrastructure/auth"
	"app/pkg/infrastructure/dcontext"
	"app/pkg/infrastructure/server/response"
	"app/pkg/interface/controller"
	"app/pkg/interface/repository"
	"app/pkg/usecase"
)

type accountHandler struct {
	AccountController controller.AccountController
}

type AccountHandler interface {
	AccountHandler(http.ResponseWriter, *http.Request)
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

func (ah *accountHandler) AccountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		log.Println("get")
		// 認証が終わった後の処理
		// contextからuserIDの取り出し
		ctx := r.Context()
		userID := dcontext.GetUserIDFromContext(ctx)
		if userID == "" {
			response.BadRequest(w, "userID is empty")
			return
		}

		res, err := ah.AccountController.ShowAccount(userID)
		if err != nil {
			response.InternalServerError(w, err.Error())
			return
		}
		// レスポンスを作成
		response.Success(w, res)

	}
}

func (ah *accountHandler) Login(w http.ResponseWriter, r *http.Request) {
	// requestbodyを読む
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	var req controller.LoginRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	// TODO: bodyのバリデーションチェック?

	// userの取得
	res, err := ah.AccountController.AuthAccount(req.UserID)

	// パスワードの認証
	err = auth.PasswordVerify(res.Password, req.Password)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	// jwtの発行
	token, err := auth.CreateToken(req.UserID)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	// クッキーに乗っける
	cookie := http.Cookie{
		Name:  "x-token",
		Value: token,
	}
	http.SetCookie(w, &cookie)

	// TODO: response
	response.Success(w, res)
}
