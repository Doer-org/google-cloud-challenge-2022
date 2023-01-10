package persistance

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/user"
	"github.com/google/uuid"
)

type UserRepository struct {
	client *ent.Client
}

func NewUserRepository(c *ent.Client) repository.IUserRepository {
	return &UserRepository{
		client: c,
	}
}

func (r *UserRepository) CreateNewUser(ctx context.Context, eu *ent.User) (*ent.User, error) {
	user, err := r.client.User.
		Create().
		SetName(eu.Name).
		SetAuthenticated(eu.Authenticated).
		SetMail(eu.Mail).
		SetIcon(eu.Icon).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("UserRepository: create user query error: %w", err)
	}
	return user, nil
}

func (r *UserRepository) GetUserById(ctx context.Context, userId uuid.UUID) (*ent.User, error) {
	user, err := r.client.User.
		Query().
		Where(user.ID(userId)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("UserRepository: get user query error: %w", err)
	}
	return user, nil
}

func (r *UserRepository) DeleteUserById(ctx context.Context, userId uuid.UUID) error {
	err := r.client.User.
		DeleteOneID(userId).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("UserRepository: delete user query error: %w", err)
	}
	return nil
}

func (r *UserRepository) UpdateUserById(ctx context.Context, userId uuid.UUID, eu *ent.User) (*ent.User, error) {
	user, err := r.client.User.
		UpdateOneID(userId).
		SetName(eu.Name).
		SetAuthenticated(eu.Authenticated).
		SetMail(eu.Mail).
		SetIcon(eu.Icon).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("UserRepository: update user query error: %w", err)
	}
	return user, nil
}

func (r *UserRepository) GetUserByMail(ctx context.Context, mail string) (*ent.User, error) {
	user, err := r.client.User.
		Query().
		Where(user.Mail(mail)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("UserRepository: get user query error: %w", err)
	}
	return user, nil
}

func (r *UserRepository) GetUserEvents(ctx context.Context, userId uuid.UUID) ([]*ent.Event, error) {
	user, err := r.client.User.
		Query().
		Where(user.ID(userId)).
		WithEvents().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: get user event query error: %w", err)
	}
	return user.Edges.Events, nil
}
