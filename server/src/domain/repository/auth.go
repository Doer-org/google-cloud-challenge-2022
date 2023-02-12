package repository

import (
	"context"
	"time"

	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type IAuth interface {
	GetTokenByUserID(ctx context.Context, userId uuid.UUID) (*oauth2.Token, error)
	StoreSession(ctx context.Context, sessionId string, userId uuid.UUID) error
	GetUserIdFromSession(ctx context.Context, sessionId string) (uuid.UUID, error)
	StoreToken(ctx context.Context, userId uuid.UUID, token *oauth2.Token) error
	UpdateToken(ctx context.Context, userId uuid.UUID, token *oauth2.Token) error
	StoreState(ctx context.Context, authState *ent.AuthStates) error
	FindStateByState(ctx context.Context, state string) (*ent.AuthStates, error)
	DeleteState(ctx context.Context, state string) error
	StoreORUpdateToken(ctx context.Context, userId uuid.UUID, token *oauth2.Token) error
	DeleteSession(ctx context.Context, sessionID string) error
	GetExpiryFromSession(ctx context.Context, sessionId string) (time.Time, error)
}
