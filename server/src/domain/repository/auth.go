package repository

import (
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type IAuth interface {
	GetTokenByUserID(userId uuid.UUID) (*oauth2.Token, error)
	StoreSession(sessionID string, userID uuid.UUID) error
	GetUserIDFromSession(sessionID string) (uuid.UUID, error)
	StoreToken(userId uuid.UUID, token *oauth2.Token) error
	UpdateToken(userId uuid.UUID, token *oauth2.Token) error
	StoreState(authState *ent.AuthStates) error
	FindStateByState(state string) (*ent.AuthStates, error)
	DeleteState(state string) error
	StoreORUpdateToken(userId uuid.UUID, token *oauth2.Token) error
}
