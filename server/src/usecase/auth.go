package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/utils"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/hash"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type Auth struct {
	repoAuth   repository.IAuth
	googleRepo repository.IGoogle
	userRepo   repository.IUser
}

func NewAuth(ra repository.IAuth, rg repository.IGoogle, ur repository.IUser) *Auth {
	return &Auth{
		repoAuth:   ra,
		googleRepo: rg,
		userRepo:   ur,
	}
}

func (uc *Auth) GetAuthURL(redirectURL string) (string, error) {
	state := hash.GetUlid()
	st := &ent.AuthStates{
		State:       state,
		RedirectURL: redirectURL,
	}
	if err := uc.repoAuth.StoreState(st); err != nil {
		return "", fmt.Errorf("StoreState: %w", err)
	}
	return uc.googleRepo.GetAuthURL(state), nil
}

func (uc *Auth) Authorization(state, code string) (string, string, error) {
	storedState, err := uc.repoAuth.FindStateByState(state)
	if err != nil {
		return "", "", fmt.Errorf("FindStateByState: %w", err)
	}
	ctx := context.Background()
	token, err := uc.googleRepo.Exchange(ctx, code)
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
	if err := uc.repoAuth.StoreSession(sessionID, userID.String()); err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("StoreSession: %w", err)
	}
	// Stateを削除するのが失敗してもログインは成功しているので、エラーを返さない
	if err := uc.repoAuth.DeleteState(state); err != nil {
		log.Printf("DeleteState: %v\n", err)
		return storedState.RedirectURL, sessionID, nil
	}
	return storedState.RedirectURL, sessionID, nil
}

// createUserIfNotExists はユーザが存在していなかったら新規に作成しIDを返します。
func (uc *Auth) createUserIfNotExists(ctx context.Context) (uuid.UUID, error) {
	user, err := uc.googleRepo.GetMe(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("GetMe: %w", err)
	}
	// uc.CreateNewUserに同じような処理があるが、ログイン時にこの関数が呼び出されるため必要
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
	gettoken, err := uc.repoAuth.GetTokenByUserID(userID)
	if err != nil && !ent.IsNotFound(err) {
		return fmt.Errorf("GetTokenByUserID: %w", err)
	}
	//TODO:消すの忘れない
	log.Println(gettoken)
	if ent.IsNotFound(err) {
		err := uc.repoAuth.StoreToken(userID, token)
		if err != nil {
			return fmt.Errorf("StoreToken: %w", err)
		}
	} else {
		err := uc.repoAuth.UpdateToken(string(userID), token)
		if err != nil {
			return fmt.Errorf("UpdateToken: %w", err)
		}
	}
	return nil
}

// GetUserIDFromSession はセッションIDから対応するユーザIDを返します。
func (uc *Auth) GetUserIDFromSession(sessionID string) (string, error) {
	userID, err := uc.repoAuth.GetUserIDFromSession(sessionID)
	if err != nil {
		return "", fmt.Errorf("GetUserIDFromSession: %w", err)
	}
	return userID, nil
}

// GetTokenByUserID は対応したユーザのアクセストークンを取得します。
func (uc *Auth) GetTokenByUserID(userID string) (*oauth2.Token, error) {
	token, err := uc.repoAuth.GetTokenByUserID(userID)
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
	newToken, err := uc.googleRepo.Refresh(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("Refresh: %w", err)
	}
	if err := uc.StoreORUpdateToken(userID, newToken); err != nil {
		return nil, fmt.Errorf("StoreORUpdateToken: %w", err)
	}
	return newToken, nil
}
