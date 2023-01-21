package persistance

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/constant"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/comment"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/event"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/user"
	"github.com/google/uuid"
)

type EventRepository struct {
	client *ent.Client
}

func NewEventRepository(c *ent.Client) repository.IEventRepository {
	return &EventRepository{
		client: c,
	}
}

func (r *EventRepository) CreateNewEvent(ctx context.Context, adminId uuid.UUID, ee *ent.Event) (*ent.Event, error) {
	event, err := r.client.Event.
		Create().
		SetName(ee.Name).
		SetDetail(ee.Detail).
		SetLocation(ee.Location).
		SetAdminID(adminId).
		SetType(string(constant.ONCE_TYPE)).
		SetState(string(constant.OPEN_STATE)).
		AddUserIDs(adminId).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("Event.Create: %w", err)
	}
	return r.getEventById(ctx, event.ID)
}

func (r *EventRepository) GetEventById(ctx context.Context, eventId uuid.UUID) (*ent.Event, error) {
	return r.getEventById(ctx, eventId)
}

func (r *EventRepository) DeleteEventById(ctx context.Context, eventId uuid.UUID) error {
	err := r.client.Event.
		DeleteOneID(eventId).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("Event.DeleteOneID: %w", err)
	}
	return nil
}

func (r *EventRepository) UpdateEventById(ctx context.Context, eventId uuid.UUID, ee *ent.Event) (*ent.Event, error) {
	event, err := r.client.Event.
		UpdateOneID(eventId).
		SetName(ee.Name).
		SetDetail(ee.Detail).
		SetLocation(ee.Location).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("Event.UpdateOneID: %w", err)
	}
	return event, nil
}

func (r *EventRepository) GetEventAdminById(ctx context.Context, eventId uuid.UUID) (*ent.User, error) {
	event, err := r.client.Event.
		Query().
		Where(event.ID(eventId)).
		WithAdmin().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("Event.Query: %w", err)
	}
	return event.Edges.Admin, nil
}

func (r *EventRepository) GetEventComments(ctx context.Context, eventId uuid.UUID) ([]*ent.Comment, error) {
	comments, err := r.client.Comment.
		Query().
		Where(comment.HasEventWith(event.ID(eventId))).
		WithUser().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("Comment.Query: %w", err)
	}
	return comments, nil
}

func (r *EventRepository) AddNewEventParticipant(ctx context.Context, eventId uuid.UUID, eu *ent.User, comment string) error {
	user, err := r.client.User.
		Create().
		SetName(eu.Name).
		SetIcon(eu.Icon).
		AddEventIDs(eventId).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("User.Create: %w", err)
	}
	if comment == "" {
		return nil
	}
	_, err = r.client.Comment.
		Create().
		SetBody(comment).
		SetEventID(eventId).
		SetUserID(user.ID).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("Comment.Create: %w", err)
	}
	return nil
}

func (r *EventRepository) ChangeEventStatusToCloseOfId(ctx context.Context, eventId uuid.UUID) (*ent.Event, error) {
	event, err := r.client.Event.
		UpdateOneID(eventId).
		SetState(string(constant.CLOSE_STATE)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("Event.UpdateOneID: %w", err)
	}
	return event, nil
}

func (r *EventRepository) ChangeEventStatusToCancelOfId(ctx context.Context, eventId uuid.UUID) (*ent.Event, error) {
	event, err := r.client.Event.
		UpdateOneID(eventId).
		SetState(string(constant.CANCEL_STATE)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("Event.UpdateOneID: %w", err)
	}
	return event, nil
}

func (r *EventRepository) GetEventUsers(ctx context.Context, eventId uuid.UUID) ([]*ent.User, error) {
	users, err := r.client.User.
		Query().
		Where(user.HasEventsWith(event.ID(eventId))).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("User.Query: %w", err)
	}
	return users, nil
}

func (r *EventRepository) getEventById(ctx context.Context, eventUuid uuid.UUID) (*ent.Event, error) {
	event, err := r.client.Event.
		Query().
		Where(event.ID(eventUuid)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("Event.Query: %w", err)
	}
	return event, nil
}
