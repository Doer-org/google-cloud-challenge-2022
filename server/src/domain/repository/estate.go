package repository

import (
	"context"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
)

type IEStateRepository interface {
	Create(ctx context.Context) (*entity.EState, error)
	UpdateStatusClose(ctx context.Context, id entity.EStateId) (*entity.EState, error)
	UpdateStatusCancel(ctx context.Context, id entity.EStateId) (*entity.EState, error)
}
