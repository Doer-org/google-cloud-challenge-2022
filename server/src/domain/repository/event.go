package repository

import (
	"context"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
)

type IEventRepository interface {
	Create(ctx context.Context, e *entity.Event, adminId entity.UserId) (*entity.Event, error)
	GetById(ctx context.Context, id entity.EventId) (*entity.Event, error)
}
