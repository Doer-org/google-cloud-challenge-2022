package persistance

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
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
		return nil, fmt.Errorf("get token by userid err : %w", err)
	}

	restoken := &oauth2.Token{}
	restoken.AccessToken = token.AccessToken
	restoken.RefreshToken = token.RefreshToken
	restoken.Expiry = token.Expiry

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

	if err != nil {
		return "", fmt.Errorf("get usedid by session err : %w", err)
	}

	return session.UserID.String(), nil
}

func (r *AuthRepository) StoreState(authState *entity.AuthState) error {
	
	return nil
}

func (r *AuthRepository) FindStateByState(state string) (*entity.AuthState, error) {

	return nil, nil
}

func (r *AuthRepository) DeleteState(state string) error {

	return nil
}
