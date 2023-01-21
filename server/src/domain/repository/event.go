package repository

import (
	"context"

	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/google/uuid"
)

type IEvent interface {
	CreateNewEvent(ctx context.Context, adminId uuid.UUID, ee *ent.Event) (*ent.Event, error)
	GetEventById(ctx context.Context, eventId uuid.UUID) (*ent.Event, error)
	DeleteEventById(ctx context.Context, eventId uuid.UUID) error
	UpdateEventById(ctx context.Context, eventId uuid.UUID, ee *ent.Event) (*ent.Event, error)
	GetEventAdminById(ctx context.Context, eventId uuid.UUID) (*ent.User, error)
	GetEventComments(ctx context.Context, eventId uuid.UUID) ([]*ent.Comment, error) //TODO: ByIdかで統一する
	AddNewEventParticipant(ctx context.Context, eventId uuid.UUID, eu *ent.User, comment string) error
	ChangeEventStatusToCloseOfId(ctx context.Context, eventId uuid.UUID) (*ent.Event, error)
	ChangeEventStatusToCancelOfId(ctx context.Context, eventId uuid.UUID) (*ent.Event, error)
	GetEventUsers(ctx context.Context, eventId uuid.UUID) ([]*ent.User, error)
}
