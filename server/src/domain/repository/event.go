package repository

import (
	"context"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
)

type IEventRepository interface {
	CreateNewEvent(ctx context.Context, e *entity.Event, adminId entity.UserId) (*entity.Event, error)
	GetEventById(ctx context.Context, eventId entity.EventId) (*entity.Event, error)
	DeleteEventById(ctx context.Context, eventId entity.EventId) error
	UpdateEventById(ctx context.Context, eventId entity.EventId, e *entity.Event) (*entity.Event, error)
	ChangeEventStatusToCloseOfId(ctx context.Context, eventId entity.EventId) (*entity.Event, error)
	ChangeEventStatusToCancelOfId(ctx context.Context, eventId entity.EventId) (*entity.Event, error)
	GetUserEvents(ctx context.Context, userId entity.UserId) ([]*entity.Event,error)
}
