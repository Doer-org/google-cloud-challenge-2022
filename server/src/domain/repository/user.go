package repository

import (
	"context"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
)

type IUserRepository interface {
	CreateNewUser(ctx context.Context, u *entity.User) (*entity.User, error)
	GetUserByMail(ctx context.Context, mail string) (*entity.User, error)
	GetUserById(ctx context.Context, id entity.UserId) (*entity.User, error)
}
