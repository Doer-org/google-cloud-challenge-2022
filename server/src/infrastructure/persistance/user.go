package persistance

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/user"
	"github.com/google/uuid"
)

type User struct {
	client *ent.Client
}

func NewUser(c *ent.Client) repository.IUser {
	return &User{
		client: c,
	}
}

func (repo *User) CreateNewUser(ctx context.Context, eu *ent.User) (*ent.User, error) {
	user, err := repo.client.User.
		Create().
		SetID(eu.ID).
		SetName(eu.Name).
		SetAuthenticated(eu.Authenticated).
		SetMail(eu.Mail).
		SetIcon(eu.Icon).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("User.Create: %w", err)
	}
	return user, nil
}

func (repo *User) GetUserById(ctx context.Context, userId uuid.UUID) (*ent.User, error) {
	user, err := repo.client.User.
		Query().
		Where(user.ID(userId)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("User.Query: %w", err)
	}
	return user, nil
}

func (repo *User) DeleteUserById(ctx context.Context, userId uuid.UUID) error {
	err := repo.client.User.
		DeleteOneID(userId).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("User.DeleteOneID: %w", err)
	}
	return nil
}

func (repo *User) UpdateUserById(ctx context.Context, userId uuid.UUID, eu *ent.User) (*ent.User, error) {
	user, err := repo.client.User.
		UpdateOneID(userId).
		SetName(eu.Name).
		SetAuthenticated(eu.Authenticated).
		SetMail(eu.Mail).
		SetIcon(eu.Icon).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("User.UpdateOneID: %w", err)
	}
	return user, nil
}

func (repo *User) GetUserByMail(ctx context.Context, mail string) (*ent.User, error) {
	user, err := repo.client.User.
		Query().
		Where(user.Mail(mail)).
		Only(ctx)
	if !ent.IsNotFound(err) && err != nil {
		return nil, fmt.Errorf("User.Query: %w", err)
	}
	return user, nil
}

func (repo *User) GetUserEvents(ctx context.Context, userId uuid.UUID) ([]*ent.Event, error) {
	user, err := repo.client.User.
		Query().
		Where(user.ID(userId)).
		WithEvents().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("User.Query: %w", err)
	}
	return user.Edges.Events, nil
}
