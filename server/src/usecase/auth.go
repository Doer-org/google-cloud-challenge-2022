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

type Auth struct {
	repo       repository.IAuth
	authGoogle google.IAuth
	userRepo   repository.IUser
}

func NewAuth(r repository.IAuth, ag google.IAuth, ur repository.IUser) *Auth {
	return &Auth{
		repo:       r,
		authGoogle: ag,
		userRepo:   ur,
	}
}

func (uc *Auth) GetAuthURL(redirectURL string) (string, error) {
	state := hash.GetUlid()
	st := &ent.AuthStates{
		State:       state,
		RedirectURL: redirectURL,
	}
	if err := uc.repo.StoreState(st); err != nil {
		return "", fmt.Errorf("StoreState: %w", err)
	}
	return uc.authGoogle.GetAuthURL(state), nil
}

func (uc *Auth) Authorization(state, code string) (string, string, error) {
	storedState, err := uc.repo.FindStateByState(state)
	if err != nil {
		return "", "", fmt.Errorf("FindStateByState: %w", err)
	}
	ctx := context.Background()
	token, err := uc.authGoogle.Exchange(ctx, code)
	if err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("Exchange: %w", err)
	}
	ctx = utils.SetTokenToContext(ctx, token)
	userID, err := uc.createUserIfNotExists(ctx)
	if err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("createUserIfNotExists: %w", err)
	}
	if err := uc.StoreORUpdateToken(userID.String(), token); err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("StoreORUpdateToken: %w", err)
	}
	sessionID := hash.GetUlid()
	if err := uc.repo.StoreSession(sessionID, userID.String()); err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("StoreSession: %w", err)
	}
	// Stateを削除するのが失敗してもログインは成功しているので、エラーを返さない
	if err := uc.repo.DeleteState(state); err != nil {
		log.Printf("DeleteState: %v\n", err)
		return storedState.RedirectURL, sessionID, nil
	}
	return storedState.RedirectURL, sessionID, nil
}

// createUserIfNotExists はユーザが存在していなかったら新規に作成しIDを返します。
func (uc *Auth) createUserIfNotExists(ctx context.Context) (uuid.UUID, error) {
	user, err := uc.authGoogle.GetMe(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("GetMe: %w", err)
	}
	res, err := uc.userRepo.GetUserByMail(ctx, user.Mail)
	if err != nil {
		return uuid.Nil, fmt.Errorf("GetUserByMail: %w", err)
	}
	if res != nil {
		return res.ID, nil
	}
	_, err = uc.userRepo.CreateNewUser(ctx, user)
	if err != nil {
		return uuid.Nil, fmt.Errorf("CreateNewUser: %w", err)
	}

	return user.ID, nil
}

func (uc *Auth) StoreORUpdateToken(userID string, token *oauth2.Token) error {
	gettoken, err := uc.repo.GetTokenByUserID(userID)
	if err != nil && !ent.IsNotFound(err) {
		return fmt.Errorf("GetTokenByUserID: %w", err)
	}
	//TODO:消すの忘れない
	log.Println(gettoken)

	if ent.IsNotFound(err) {
		err := uc.repo.StoreToken(userID, token)
		if err != nil {
			return fmt.Errorf("StoreToken: %w", err)
		}

	} else {
		err := uc.repo.UpdateToken(string(userID), token)
		if err != nil {
			return fmt.Errorf("UpdateToken: %w", err)
		}

	}
	return nil
}

// GetUserIDFromSession はセッションIDから対応するユーザIDを返します。
func (uc *Auth) GetUserIDFromSession(sessionID string) (string, error) {
	userID, err := uc.repo.GetUserIDFromSession(sessionID)
	if err != nil {
		return "", fmt.Errorf("GetUserIDFromSession: %w", err)
	}
	return userID, nil
}

// GetTokenByUserID は対応したユーザのアクセストークンを取得します。
func (uc *Auth) GetTokenByUserID(userID string) (*oauth2.Token, error) {
	token, err := uc.repo.GetTokenByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("GetTokenByUserID: %w", err)
	}
	return token, nil
}

// RefreshAccessToken はリフレッシュトークンを使用してアクセストークンを更新し保存します。
func (uc *Auth) RefreshAccessToken(userID string, token *oauth2.Token) (*oauth2.Token, error) {
	if token.Valid() {
		return token, nil
	}
	ctx := context.Background()
	newToken, err := uc.authGoogle.Refresh(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("Refresh: %w", err)
	}
	if err := uc.StoreORUpdateToken(userID, newToken); err != nil {
		return nil, fmt.Errorf("StoreORUpdateToken: %w", err)
	}
	return newToken, nil
}
