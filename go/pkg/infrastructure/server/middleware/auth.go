package middleware

import (
	"app/pkg/domain"
	"app/pkg/infrastructure/auth"
	"app/pkg/infrastructure/dcontext"
	"app/pkg/infrastructure/server/response"
	"context"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

func Authorized(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if ctx == nil {
			ctx = context.Background()
		}

		// cookieからtokenを取得
		cookie, err := r.Cookie("x-token")
		if err != nil {
			response.HttpError(w, domain.Unauthorized(err))
			return
		}

		// tokenの検証
		token, err := auth.VerifyToken(cookie.Value)
		if err != nil {
			response.HttpError(w, domain.Unauthorized(err))
			return
		}

		// tokenからuserIDの取り出し
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["id"].(string)

		// TODO: ユーザが存在するか確認. DBにアクセスかける

		// contextにuserIDを保存
		ctx = dcontext.SetUserID(ctx, userID)

		// nextfnc
		nextFunc(w, r.WithContext(ctx))
	}
}
