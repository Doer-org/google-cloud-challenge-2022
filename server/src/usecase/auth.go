package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	mycontext "github.com/Doer-org/google-cloud-challenge-2022/utils/context"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/hash"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type Auth struct {
	repoAuth   repository.IAuth
	googleRepo repository.IGoogle
	userRepo   repository.IUser
}

type IAuth interface {
	GetAuthURL(ctx context.Context, redirectURL string) (url string, state string, err error)
	Authorization(ctx context.Context, state, code string) (string, string, error)
	GetUserIdFromSession(ctx context.Context, sessionId string) (uuid.UUID, error)
	GetTokenByUserId(ctx context.Context, userId uuid.UUID) (*oauth2.Token, error)
	RefreshAccessToken(ctx context.Context, userId uuid.UUID, token *oauth2.Token) (*oauth2.Token, error)
}

func NewAuth(ra repository.IAuth, rg repository.IGoogle, ur repository.IUser) IAuth {
	return &Auth{
		repoAuth:   ra,
		googleRepo: rg,
		userRepo:   ur,
	}
}

func (uc *Auth) GetAuthURL(ctx context.Context, redirectURL string) (url string, resstate string, err error) {
	state := hash.GetUlid()
	st := &ent.AuthStates{
		State:       state,
		RedirectURL: redirectURL,
	}
	if err := uc.repoAuth.StoreState(ctx,st); err != nil {
		return "", "", fmt.Errorf("storeState: %w", err)
	}

	return uc.googleRepo.GetAuthURL(state), state, nil
}

// memo: 複数のブラウザを立ち上げた場合、sessionが複数作られる
func (uc *Auth) Authorization(ctx context.Context, state, code string) (string, string, error) {
	storedState, err := uc.repoAuth.FindStateByState(ctx,state)
	if err != nil {
		return "", "", fmt.Errorf("findStateByState: %w", err)
	}
	token, err := uc.googleRepo.Exchange(ctx, code)
	if err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("exchange: %w", err)
	}
	ctx = mycontext.SetToken(ctx, token)
	userId, err := uc.createUserIfNotExists(ctx)
	if err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("createUserIfNotExists: %w", err)
	}
	if err := uc.repoAuth.StoreORUpdateToken(ctx, userId, token); err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("storeORUpdateToken: %w", err)
	}
	sessionID := hash.GetUlid()
	if err := uc.repoAuth.StoreSession(ctx, sessionID, userId); err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("storeSession: %w", err)
	}
	// Stateを削除するのが失敗してもログインは成功しているので、エラーを返さない
	//TODO: stateはなんのために使われてるのか?
	if err := uc.repoAuth.DeleteState(ctx, state); err != nil {
		log.Println("DeleteState: %w", err)
		return storedState.RedirectURL, sessionID, nil
	}
	return storedState.RedirectURL, sessionID, nil
}

// createUserIfNotExists はユーザが存在していなかったら新規に作成しIDを返します。
func (uc *Auth) createUserIfNotExists(ctx context.Context) (uuid.UUID, error) {
	user, err := uc.googleRepo.GetMe(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("getMe: %w", err)
	}
	userId, found := uc.checkUserExistsByMail(ctx, user.Mail)
	if found {
		return userId, nil
	}
	newUser, err := uc.userRepo.CreateNewUser(ctx, user)
	if err != nil {
		return uuid.Nil, fmt.Errorf("createNewUser: %w", err)
	}
	return newUser.ID, nil
}

func (uc *Auth) checkUserExistsByMail(ctx context.Context, mail string) (uuid.UUID, bool) {
	found, _ := uc.userRepo.GetUserByMail(ctx, mail)
	if found != nil {
		return found.ID, true
	}
	return uuid.Nil, false
}

// GetUserIDFromSession はセッションIDから対応するユーザIDを返します。
func (uc *Auth) GetUserIdFromSession(ctx context.Context, sessionId string) (uuid.UUID, error) {
	userId, err := uc.repoAuth.GetUserIdFromSession(ctx,sessionId)
	if err != nil {
		return uuid.Nil, fmt.Errorf("getUserIDFromSession: %w", err)
	}
	return userId, nil
}

// GetTokenByUserID は対応したユーザのアクセストークンを取得します。
func (uc *Auth) GetTokenByUserId(ctx context.Context, userId uuid.UUID) (*oauth2.Token, error) {
	token, err := uc.repoAuth.GetTokenByUserID(ctx,userId)
	if err != nil {
		return nil, fmt.Errorf("getTokenByUserID: %w", err)
	}
	return token, nil
}

// RefreshAccessToken はリフレッシュトークンを使用してアクセストークンを更新し保存します。
func (uc *Auth) RefreshAccessToken(ctx context.Context, userId uuid.UUID, token *oauth2.Token) (*oauth2.Token, error) {
	if token.Valid() {
		return token, nil
	}
	newToken, err := uc.googleRepo.Refresh(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("refresh: %w", err)
	}
	if err := uc.repoAuth.StoreORUpdateToken(ctx, userId, newToken); err != nil {
		return nil, fmt.Errorf("storeORUpdateToken: %w", err)
	}
	return newToken, nil
}
