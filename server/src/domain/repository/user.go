package repository

import (
	"context"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
)

type IUserRepository interface {
	Create(ctx context.Context, u *entity.User) error
	GetByMail(ctx context.Context, mail string) (*entity.User, error)
	GetById(ctx context.Context, id entity.UserId) (*entity.User, error)
}
