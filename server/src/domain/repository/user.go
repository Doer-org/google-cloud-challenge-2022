package repository

import (
	"context"

	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/google/uuid"
)

type IUserRepository interface {
	CreateNewUser(ctx context.Context, eu *ent.User) (*ent.User, error)
	GetUserById(ctx context.Context, userId uuid.UUID) (*ent.User, error)
	DeleteUserById(ctx context.Context, userId uuid.UUID) error
	UpdateUserById(ctx context.Context, userId uuid.UUID, eu *ent.User) (*ent.User, error)
	GetUserByMail(ctx context.Context, mail string) (*ent.User, error)
	GetUserEvents(ctx context.Context, userId uuid.UUID) ([]*ent.Event, error)
}
