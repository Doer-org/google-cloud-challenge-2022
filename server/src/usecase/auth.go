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
		return "", fmt.Errorf("storeState: %w", err)
	}
	return uc.googleRepo.GetAuthURL(state), nil
}

// TODO: interfaceで統一すべき
// memo: 複数のブラウザを立ち上げた場合、sessionが複数作られる
func (uc *Auth) Authorization(state, code string) (string, string, error) {
	storedState, err := uc.repoAuth.FindStateByState(state)
	if err != nil {
		return "", "", fmt.Errorf("findStateByState: %w", err)
	}
	ctx := context.Background()
	token, err := uc.googleRepo.Exchange(ctx, code)
	if err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("exchange: %w", err)
	}
	ctx = mycontext.SetToken(ctx, token)
	userId, err := uc.createUserIfNotExists(ctx)
	if err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("createUserIfNotExists: %w", err)
	}
	// TODO: contextを引数に追加
	if err := uc.repoAuth.StoreORUpdateToken(userId, token); err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("storeORUpdateToken: %w", err)
	}
	sessionID := hash.GetUlid()
	//TODO: contextを引数に追加
	if err := uc.repoAuth.StoreSession(sessionID, userId); err != nil {
		return storedState.RedirectURL, "", fmt.Errorf("storeSession: %w", err)
	}
	// Stateを削除するのが失敗してもログインは成功しているので、エラーを返さない
	//TODO: stateはなんのために使われるんだろう..
	if err := uc.repoAuth.DeleteState(state); err != nil {
		log.Println("DeleteState: %v", err)
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
	userId,found := uc.checkUserExistsByMail(ctx,user.Mail)
	if found {
		return userId,nil
	}
	newUser,err := uc.userRepo.CreateNewUser(ctx, user)
	if err != nil {
		return uuid.Nil, fmt.Errorf("createNewUser: %w", err)
	}
	return newUser.ID,nil
}

func (uc *Auth) checkUserExistsByMail(ctx context.Context, mail string) (uuid.UUID,bool) {
	found, _ := uc.userRepo.GetUserByMail(ctx,  mail)
	if found != nil {
		return found.ID,true
	}
	return uuid.Nil,false
}

// GetUserIDFromSession はセッションIDから対応するユーザIDを返します。
func (uc *Auth) GetUserIdFromSession(sessionId string) (uuid.UUID, error) {
	userId, err := uc.repoAuth.GetUserIdFromSession(sessionId)
	if err != nil {
		return uuid.Nil, fmt.Errorf("getUserIDFromSession: %w", err)
	}
	return userId, nil
}

// GetTokenByUserID は対応したユーザのアクセストークンを取得します。
func (uc *Auth) GetTokenByUserId(userId uuid.UUID) (*oauth2.Token, error) {
	token, err := uc.repoAuth.GetTokenByUserID(userId)
	if err != nil {
		return nil, fmt.Errorf("getTokenByUserID: %w", err)
	}
	return token, nil
}

// RefreshAccessToken はリフレッシュトークンを使用してアクセストークンを更新し保存します。
func (uc *Auth) RefreshAccessToken(userId uuid.UUID, token *oauth2.Token) (*oauth2.Token, error) {
	if token.Valid() {
		return token, nil
	}
	ctx := context.Background()
	newToken, err := uc.googleRepo.Refresh(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("refresh: %w", err)
	}
	if err := uc.repoAuth.StoreORUpdateToken(userId, newToken); err != nil {
		return nil, fmt.Errorf("storeORUpdateToken: %w", err)
	}
	return newToken, nil
}
