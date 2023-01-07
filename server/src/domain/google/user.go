package google

import (
	"context"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
)

type IUser interface {
	GetMe(ctx context.Context) (*entity.User, error)
}
