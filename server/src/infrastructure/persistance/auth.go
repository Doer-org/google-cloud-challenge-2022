package persistance

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/authstates"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/googleauth"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/loginsessions"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type Auth struct {
	Client *ent.Client
}

func NewAuth(c *ent.Client) repository.IAuth {
	return &Auth{
		Client: c,
	}
}

func (repo *Auth) StoreToken(userId uuid.UUID, token *oauth2.Token) error {
	_, err := repo.Client.GoogleAuth.
		Create().
		SetAccessToken(token.AccessToken).
		SetRefreshToken(token.RefreshToken).
		SetExpiry(token.Expiry).
		SetUserID(userId).
		Save(context.Background())
	if err != nil {
		return fmt.Errorf("googleAuth.Create: %w", err)
	}
	return nil
}

func (repo *Auth) UpdateToken(userId uuid.UUID, token *oauth2.Token) error {
	_, err := repo.Client.GoogleAuth.Update().
		SetAccessToken(token.AccessToken).
		SetRefreshToken(token.RefreshToken).
		SetExpiry(token.Expiry).
		Where(googleauth.UserID(userId)).
		Save(context.Background())
	if err != nil {
		return fmt.Errorf("googleAuth.Update: %w", err)
	}
	return nil
}

// もしtokenが存在した場合,更新し, 存在しなかった場合新たにstoreする
func (repo *Auth) StoreORUpdateToken(userId uuid.UUID, token *oauth2.Token) error {
	found, err := repo.GetTokenByUserID(userId)
	if err != nil && !ent.IsNotFound(err) {
		return fmt.Errorf("getTokenByUserID: %w", err)
	}
	if found != nil {
		if err := repo.UpdateToken(userId, token); err != nil {
			return fmt.Errorf("updateToken: %w", err)
		}
		return nil
	} else if ent.IsNotFound(err) {
		if err := repo.StoreToken(userId, token); err != nil {
			return fmt.Errorf("storeToken: %w", err)
		}
	} else {
		return fmt.Errorf("found is empty, and ent is not found")
	}
	return nil
}

func (repo *Auth) GetTokenByUserID(userId uuid.UUID) (*oauth2.Token, error) {
	token, err := repo.Client.GoogleAuth.
		Query().
		Where(googleauth.UserID(userId)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}
	return &oauth2.Token{
		AccessToken: token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry: token.Expiry,
	},nil
}

func (repo *Auth) StoreSession(sessionID string, userId uuid.UUID) error {
	_, err := repo.Client.LoginSessions.
		Create().
		SetUserID(userId).
		SetID(sessionID).
		Save(context.Background())
	if err != nil {
		return fmt.Errorf("loginSessions.Create: %w", err)
	}
	return nil
}

func (repo *Auth) GetUserIdFromSession(sessionId string) (uuid.UUID, error) {
	session, err := repo.Client.LoginSessions.
		Query().
		Where(loginsessions.ID(sessionId)).
		Only(context.Background())
	if err != nil && !ent.IsNotFound(err) {
		return uuid.Nil, fmt.Errorf("loginSessions.Query: %w", err)
	}
	return session.UserID, nil
}

func (repo *Auth) StoreState(authState *ent.AuthStates) error {
	_, err := repo.Client.AuthStates.
		Create().
		SetState(authState.State).
		SetRedirectURL(authState.RedirectURL).
		Save(context.Background())
	if err != nil {
		return fmt.Errorf("authStates.Create: %w", err)
	}
	return nil
}

func (repo *Auth) FindStateByState(state string) (*ent.AuthStates, error) {
	resstate, err := repo.Client.AuthStates.
		Query().
		Where(authstates.State(state)).
		Only(context.Background())
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("authStates.Query: %w", err)
	}
	res := &ent.AuthStates{}
	res.RedirectURL = resstate.RedirectURL
	res.State = resstate.State
	return res, nil
}

func (repo *Auth) DeleteState(state string) error {
	_, err := repo.Client.AuthStates.
		Delete().
		Where(authstates.State(state)).
		Exec(context.Background())
	if err != nil {
		return fmt.Errorf("authStates.Delete: %w", err)
	}
	return nil
}
