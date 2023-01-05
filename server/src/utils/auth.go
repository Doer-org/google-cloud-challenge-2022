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
