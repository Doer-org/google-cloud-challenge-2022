package repository

import (
	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"golang.org/x/oauth2"
)

type IAuthRepository interface {
	GetTokenByUserID(userID string) (*oauth2.Token, error)
	StoreSession(sessionID, userID string) error
	GetUserIDFromSession(sessionID string) (string, error)
	StoreToken(userId string, token *oauth2.Token) error
	UpdateToken(userId string, token *oauth2.Token) error

	StoreState(authState *entity.AuthState) error
	FindStateByState(state string) (*entity.AuthState, error)
	DeleteState(state string) error
}
