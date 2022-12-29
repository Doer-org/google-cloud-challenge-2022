package repository

import (
	"context"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
)

type IUserRepository interface {
	Create(ctx context.Context,u *entity.User) error
}
