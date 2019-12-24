package dcontext

import (
	"context"
)

type key string

const (
	userIDKey key = "userID"
)

// SetUserID userIDをcontextに保存
func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

// GetUserIDFromContext contextからユーザIDを取得
func GetUserIDFromContext(ctx context.Context) string {
	var userID string
	if ctx.Value(userIDKey) != nil {
		userID = ctx.Value(userIDKey).(string)
	}
	return userID

}
