package usecase

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/google"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/hash"
	"golang.org/x/oauth2"
)

type IAuthUsecase interface {
	GetTokenByUserID(userID string) (*oauth2.Token, error)
	StoreSession(sessionID, userID string) error
	GetUserIDFromSession(sessionID string) (string, error)

	StoreState(authState *entity.AuthState) error
	FindStateByState(state string) (*entity.AuthState, error)
	DeleteState(state string) error
}

type AuthUsecase struct {
	repo       repository.IAuthRepository
	authGoogle google.Auth
	userRepo   repository.IUserRepository
}

func NewAuthrUsecase(r repository.IAuthRepository) IAuthUsecase {
	return &AuthUsecase{
		repo: r,
	}
}

func (u *AuthUsecase) GetAuthURL(redirectURL string) (string, error) {
	state := hash.GetUlid()
	st := &entity.AuthState{
		State:       state,
		RedirectURL: redirectURL,
	}

	if err := u.repo.StoreState(st); err != nil {
		return "", fmt.Errorf("store state for authorization: %w", err)
	}
	return u.authGoogle.GetAuthURL(state), nil
}

func (u *AuthUsecase) Authorization(state, code string) (string, string, error) {
	storedState, err := u.repo.FindStateByState(state)
	if err != nil {
		return "", "", fmt.Errorf("find temp state state=%s: %w", state, err)
	}

	ctx := context.Background()
	token, err := u.authGoogle.Exchange(ctx, code)
	if err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("exchange and get oauth2 token: %w", err)
	}

}

func (u *AuthUsecase) createUserIfNotExists(ctx context.Context) (string, error) {
	
}
