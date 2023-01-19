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

type AuthRepository struct {
	Client *ent.Client
}

func NewAuthRepository(c *ent.Client) repository.IAuthRepository {
	return &AuthRepository{
		Client: c,
	}
}

func (r *AuthRepository) StoreToken(userId string, token *oauth2.Token) error {
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		return fmt.Errorf("uuid parse err : %w", err)
	}

	_, err = r.Client.GoogleAuth.
		Create().
		SetAccessToken(token.AccessToken).
		SetRefreshToken(token.RefreshToken).
		SetExpiry(token.Expiry).
		SetUserID(userUuid).
		Save(context.Background())

	if err != nil {
		return fmt.Errorf("create token err : %w", err)
	}

	return nil
}

func (r *AuthRepository) UpdateToken(userId string, token *oauth2.Token) error {
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		return fmt.Errorf("uuid parse err : %w", err)
	}

	_, err = r.Client.GoogleAuth.Update().
		SetAccessToken(token.AccessToken).
		SetRefreshToken(token.RefreshToken).
		SetExpiry(token.Expiry).
		Where(googleauth.UserID(userUuid)).
		Save(context.Background())

	if err != nil {
		return fmt.Errorf("update token err : %w", err)
	}

	return nil
}

func (r *AuthRepository) GetTokenByUserID(userID string) (*oauth2.Token, error) {
	userUuid, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("uuid parse err : %w", err)
	}

	token, err := r.Client.GoogleAuth.
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

func (r *AuthRepository) StoreSession(sessionID, userID string) error {
	userUuid, err := uuid.Parse(userID)
	if err != nil {
		return fmt.Errorf("uuid parse err : %w", err)
	}

	_, err = r.Client.LoginSessions.
		Create().
		SetUserID(userUuid).
		SetID(sessionID).
		Save(context.Background())

	if err != nil {
		return fmt.Errorf("create session err : %w", err)
	}
	return nil
}

func (r *AuthRepository) GetUserIDFromSession(sessionID string) (string, error) {
	session, err := r.Client.LoginSessions.
		Query().
		Where(loginsessions.ID(sessionID)).
		Only(context.Background())

	if err != nil && !ent.IsNotFound(err) {
		return "", fmt.Errorf("get usedid by session err : %w", err)
	}

	return session.UserID.String(), nil
}

func (r *AuthRepository) StoreState(authState *ent.AuthStates) error {
	_, err := r.Client.AuthStates.
		Create().
		SetState(authState.State).
		SetRedirectURL(authState.RedirectURL).
		Save(context.Background())

	if err != nil {
		return fmt.Errorf("create state err : %w", err)
	}

	return nil
}

func (r *AuthRepository) FindStateByState(state string) (*ent.AuthStates, error) {
	resstate, err := r.Client.AuthStates.
		Query().
		Where(authstates.State(state)).
		Only(context.Background())

	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("get state by state err : %w", err)
	}

	res := &ent.AuthStates{}
	res.RedirectURL = resstate.RedirectURL
	res.State = resstate.State

	return res, nil
}

func (r *AuthRepository) DeleteState(state string) error {
	_, err := r.Client.AuthStates.
		Delete().
		Where(authstates.State(state)).
		Exec(context.Background())

	if err != nil {
		return fmt.Errorf("delete state by state err : %w", err)
	}

	return nil
}
