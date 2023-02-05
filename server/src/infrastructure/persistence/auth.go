package persistence

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/oauth2"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/authstates"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/googleauth"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/loginsessions"
)

type Auth struct {
	client *ent.Client
}

func NewAuth(c *ent.Client) repository.IAuth {
	return &Auth{
		client: c,
	}
}

func (repo *Auth) StoreToken(ctx context.Context, userId uuid.UUID, token *oauth2.Token) error {
	_, err := repo.client.GoogleAuth.
		Create().
		SetAccessToken(token.AccessToken).
		SetRefreshToken(token.RefreshToken).
		SetExpiry(token.Expiry).
		SetUserID(userId).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("googleAuth.Create: %w", err)
	}
	return nil
}

func (repo *Auth) UpdateToken(ctx context.Context, userId uuid.UUID, token *oauth2.Token) error {
	_, err := repo.client.GoogleAuth.Update().
		SetAccessToken(token.AccessToken).
		SetRefreshToken(token.RefreshToken).
		SetExpiry(token.Expiry).
		Where(googleauth.UserID(userId)).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("googleAuth.Update: %w", err)
	}
	return nil
}

// もしtokenが存在した場合,更新し, 存在しなかった場合新たにstoreする
func (repo *Auth) StoreORUpdateToken(ctx context.Context, userId uuid.UUID, token *oauth2.Token) error {
	found, err := repo.GetTokenByUserID(ctx, userId)
	if err != nil && !ent.IsNotFound(err) {
		return fmt.Errorf("getTokenByUserID: %w", err)
	}
	if found != nil {
		if err := repo.UpdateToken(ctx, userId, token); err != nil {
			return fmt.Errorf("updateToken: %w", err)
		}
		return nil
	} else if ent.IsNotFound(err) {
		if err := repo.StoreToken(ctx, userId, token); err != nil {
			return fmt.Errorf("storeToken: %w", err)
		}
	} else {
		return fmt.Errorf("found is empty, and ent is not found")
	}
	return nil
}

func (repo *Auth) GetTokenByUserID(ctx context.Context, userId uuid.UUID) (*oauth2.Token, error) {
	token, err := repo.client.GoogleAuth.
		Query().
		Where(googleauth.UserID(userId)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return &oauth2.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}, nil
}

func (repo *Auth) StoreSession(ctx context.Context, sessionID string, userId uuid.UUID) error {
	_, err := repo.client.LoginSessions.
		Create().
		SetUserID(userId).
		SetID(sessionID).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("loginSessions.Create: %w", err)
	}
	return nil
}

func (repo *Auth) GetUserIdFromSession(ctx context.Context, sessionId string) (uuid.UUID, error) {
	session, err := repo.client.LoginSessions.
		Query().
		Where(loginsessions.ID(sessionId)).
		Only(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("loginSessions.Query: %w", err)
	}
	return session.UserID, nil
}

func (repo *Auth) StoreState(ctx context.Context, authState *ent.AuthStates) error {
	_, err := repo.client.AuthStates.
		Create().
		SetState(authState.State).
		SetRedirectURL(authState.RedirectURL).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("authStates.Create: %w", err)
	}
	return nil
}

func (repo *Auth) FindStateByState(ctx context.Context, state string) (*ent.AuthStates, error) {
	authState, err := repo.client.AuthStates.
		Query().
		Where(authstates.State(state)).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("authStates.Query: %w", err)
	}
	return authState, nil
}

func (repo *Auth) DeleteState(ctx context.Context, state string) error {
	_, err := repo.client.AuthStates.
		Delete().
		Where(authstates.State(state)).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("authStates.Delete: %w", err)
	}
	return nil
}
