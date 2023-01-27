package context

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type ContextKey string

var (
	userIdKey    ContextKey = "userIdKey"
	creatorIdKey ContextKey = "creatorIdKey"
	tokenKey     ContextKey = "tokenKey"
)

func SetToken(ctx context.Context, token *oauth2.Token) context.Context {
	if token != nil {
		return context.WithValue(ctx, tokenKey, token)
	}
	return ctx
}

func GetToken(ctx context.Context) (*oauth2.Token, bool) {
	v := ctx.Value(tokenKey)
	token, ok := v.(*oauth2.Token)
	return token, ok
}

func SetUserId(ctx context.Context, userId string) context.Context {
	if userId != "" {
		return context.WithValue(ctx, userIdKey, userId)
	}
	return ctx
}

func GetUser(ctx context.Context) (uuid.UUID,bool) {
	v := ctx.Value(userIdKey)
	userIdString,ok := v.(string)
	if !ok {
		return uuid.Nil,ok
	}
	userId,err := uuid.Parse(userIdString)
	if err != nil {
		log.Println("error: getUser from context failed")
		return uuid.Nil,false
	}
	return userId,ok
}
// ctxに入っているsessionのuserIdと引数で受け取ったuserIdを比較します
func CompareUserIdAndUserSessionId(ctx context.Context,userId uuid.UUID) error {
	userSessId,ok := GetUser(ctx)
	if !ok {
		return fmt.Errorf("GetUser: failed to get user from context")
	}
	if userId == uuid.Nil {
		return fmt.Errorf("func args userId is nil")
	}
	if userId.String() != userSessId.String() {
		return fmt.Errorf("userId is not matched session userId")
	}
	return nil
}