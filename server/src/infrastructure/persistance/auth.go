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

func (repo *Auth) StoreToken(userId string, token *oauth2.Token) error {
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		return fmt.Errorf("uuid.Parse: %w", err)
	}
	_, err = repo.Client.GoogleAuth.
		Create().
		SetAccessToken(token.AccessToken).
		SetRefreshToken(token.RefreshToken).
		SetExpiry(token.Expiry).
		SetUserID(userUuid).
		Save(context.Background())
	if err != nil {
		return fmt.Errorf("GoogleAuth.Create: %w", err)
	}
	return nil
}

func (repo *Auth) UpdateToken(userId string, token *oauth2.Token) error {
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		return fmt.Errorf("uuid.Parse: %w", err)
	}
	_, err = repo.Client.GoogleAuth.Update().
		SetAccessToken(token.AccessToken).
		SetRefreshToken(token.RefreshToken).
		SetExpiry(token.Expiry).
		Where(googleauth.UserID(userUuid)).
		Save(context.Background())
	if err != nil {
		return fmt.Errorf("GoogleAuth.Update: %w", err)
	}
	return nil
}

func (repo *Auth) GetTokenByUserID(userID string) (*oauth2.Token, error) {
	userUuid, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse: %w", err)
	}
	token, err := repo.Client.GoogleAuth.
		Query().
		Where(googleauth.UserID(userUuid)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}
	restoken := &oauth2.Token{}
	if token != nil {
		restoken.AccessToken = token.AccessToken
		restoken.RefreshToken = token.RefreshToken
		restoken.Expiry = token.Expiry
	}
	return restoken, nil
}

func (repo *Auth) StoreSession(sessionID, userID string) error {
	userUuid, err := uuid.Parse(userID)
	if err != nil {
		return fmt.Errorf("uuid.Parse: %w", err)
	}
	_, err = repo.Client.LoginSessions.
		Create().
		SetUserID(userUuid).
		SetID(sessionID).
		Save(context.Background())
	if err != nil {
		return fmt.Errorf("LoginSessions.Create: %w", err)
	}
	return nil
}

func (repo *Auth) GetUserIDFromSession(sessionID string) (string, error) {
	session, err := repo.Client.LoginSessions.
		Query().
		Where(loginsessions.ID(sessionID)).
		Only(context.Background())
	if err != nil && !ent.IsNotFound(err) {
		return "", fmt.Errorf("LoginSessions.Query: %w", err)
	}
	return session.UserID.String(), nil
}

func (repo *Auth) StoreState(authState *ent.AuthStates) error {
	_, err := repo.Client.AuthStates.
		Create().
		SetState(authState.State).
		SetRedirectURL(authState.RedirectURL).
		Save(context.Background())
	if err != nil {
		return fmt.Errorf("AuthStates.Create: %w", err)
	}
	return nil
}

func (repo *Auth) FindStateByState(state string) (*ent.AuthStates, error) {
	resstate, err := repo.Client.AuthStates.
		Query().
		Where(authstates.State(state)).
		Only(context.Background())
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("AuthStates.Query: %w", err)
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
		return fmt.Errorf("AuthStates.Delete: %w", err)
	}
	return nil
}
