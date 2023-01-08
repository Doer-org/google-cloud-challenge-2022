package utils

import (
	"context"

	"golang.org/x/oauth2"
)

type ContextKey string

var (
	userIDKey    ContextKey = "userIDKey"
	creatorIDKey ContextKey = "creatorIDKey"
	tokenKey     ContextKey = "tokenKey"
)

func SetTokenToContext(ctx context.Context, token *oauth2.Token) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

func GetTokenFromContext(ctx context.Context) (*oauth2.Token, bool) {
	v := ctx.Value(tokenKey)
	token, ok := v.(*oauth2.Token)
	return token, ok
}

func SetUserIDToContext(ctx context.Context, userID string) context.Context {
	if userID != "" {
		return context.WithValue(ctx, userIDKey, userID)
	}
	return ctx
}

