package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type key string

const (
	secretKey key = "note_app"
)

// CreateToken JWTの発行
func CreateToken(userID string) (string, error) {
	// tokenの作成
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	// claimsの設定
	token.Claims = jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	// 署名
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

// VerifyToken JWTの検証
func VerifyToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil

}
