package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/google"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/utils"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/hash"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type AuthUsecase struct {
	repo       repository.IAuthRepository
	authGoogle google.IAuth

	userRepo repository.IUserRepository
}

func NewAuthUsecase(r repository.IAuthRepository, ag google.IAuth, ur repository.IUserRepository) *AuthUsecase {
	return &AuthUsecase{
		repo:       r,
		authGoogle: ag,
		userRepo:   ur,
	}
}

func (u *AuthUsecase) GetAuthURL(redirectURL string) (string, error) {
	state := hash.GetUlid()
	st := &ent.AuthStates{
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

	ctx = utils.SetTokenToContext(ctx, token)
	userID, err := u.createUserIfNotExists(ctx)
	if err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("get or create user: %w", err)
	}

	if err := u.StoreORUpdateToken(userID.String(), token); err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("store or update oauth token though repo userID=%s: %w", userID, err)
	}

	sessionID := hash.GetUlid()
	if err := u.repo.StoreSession(sessionID, userID.String()); err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("store session sessionID=%s userID=%s : %w", sessionID, userID, err)
	}

	// Stateを削除するのが失敗してもログインは成功しているので、エラーを返さない
	if err := u.repo.DeleteState(state); err != nil {
		log.Printf("Failed to delete state state=%s: %v\n", state, err)
		return storedState.RedirectURL, sessionID, nil
	}

	return storedState.RedirectURL, sessionID, nil
}

// createUserIfNotExists はユーザが存在していなかったら新規に作成しIDを返します。
func (u *AuthUsecase) createUserIfNotExists(ctx context.Context) (uuid.UUID, error) {
	user, err := u.authGoogle.GetMe(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("get my info from Google: %w", err)
	}

	getuser, err := u.userRepo.GetUserById(ctx, user.ID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("get user by id: %w", err)
	}

	if getuser.Name != "" {
		userId := getuser.ID
		return userId, nil
	}

	_, err = u.userRepo.CreateNewUser(ctx, user)
	if err != nil {
		return uuid.Nil, fmt.Errorf("create user by id: %w", err)
	}

	return user.ID, nil
}

func (u *AuthUsecase) StoreORUpdateToken(userID string, token *oauth2.Token) error {
	gettoken, err := u.repo.GetTokenByUserID(userID)
	if err != nil {
		return fmt.Errorf("get token from userId userID=%s: %w", userID, err)
	}

	if gettoken.AccessToken == "" && gettoken.RefreshToken == "" {
		err := u.repo.StoreToken(string(userID), token)
		if err != nil {
			return fmt.Errorf("create token from userId userID=%s: %w", userID, err)
		}
	} else {
		err := u.repo.UpdateToken(string(userID), token)
		if err != nil {
			return fmt.Errorf("update token from userId userID=%s: %w", userID, err)
		}
	}

	return nil
}

// GetUserIDFromSession はセッションIDから対応するユーザIDを返します。
func (u *AuthUsecase) GetUserIDFromSession(sessionID string) (string, error) {
	userID, err := u.repo.GetUserIDFromSession(sessionID)
	if err != nil {
		return "", fmt.Errorf("get user from session sessionID=%s: %w", sessionID, err)
	}
	return userID, nil
}

// GetTokenByUserID は対応したユーザのアクセストークンを取得します。
func (u *AuthUsecase) GetTokenByUserID(userID string) (*oauth2.Token, error) {
	token, err := u.repo.GetTokenByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("get oauth token userID=%s: %w", userID, err)
	}
	return token, nil
}

// RefreshAccessToken はリフレッシュトークンを使用してアクセストークンを更新し保存します。
func (u *AuthUsecase) RefreshAccessToken(userID string, token *oauth2.Token) (*oauth2.Token, error) {
	if token.Valid() {
		return token, nil
	}
	ctx := context.Background()
	newToken, err := u.authGoogle.Refresh(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("refresh access token through spotify client: %w", err)
	}

	if err := u.StoreORUpdateToken(userID, newToken); err != nil {
		return nil, fmt.Errorf("update new token: %w", err)
	}
	return newToken, nil
}
