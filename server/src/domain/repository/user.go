package repository

import (
	"context"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
)

type IUserRepository interface {
	CreateNewUser(ctx context.Context, u *entity.User) (*entity.User, error)
	GetUserById(ctx context.Context, userId entity.UserId) (*entity.User, error)
	DeleteUserById(ctx context.Context, userId entity.UserId) error
	UpdateUserById(ctx context.Context, userId entity.UserId, u *entity.User) (*entity.User, error)
	GetUserByMail(ctx context.Context, mail string) (*entity.User, error)
	GetEventAdminById(ctx context.Context,eventId entity.EventId) (*entity.User, error)
}
