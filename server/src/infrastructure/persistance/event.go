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

type Event struct {
	client *ent.Client
}

func NewEvent(c *ent.Client) repository.IEvent {
	return &Event{
		client: c,
	}
}

func (repo *Event) CreateNewEvent(ctx context.Context, adminId uuid.UUID, ee *ent.Event) (*ent.Event, error) {
	event, err := repo.client.Event.
		Create().
		SetName(ee.Name).
		SetDetail(ee.Detail).
		SetLocation(ee.Location).
		SetSize(ee.Size).
		SetLimitHour(ee.LimitHour).
		SetAdminID(adminId).
		SetType(string(constant.TYPE_ONCE)).
		SetState(string(constant.STATE_OPEN)).
		AddUserIDs(adminId).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("event.Create: %w", err)
	}
	return repo.getEventById(ctx, event.ID)
}

func (repo *Event) GetEventById(ctx context.Context, eventId uuid.UUID) (*ent.Event, error) {
	return repo.getEventById(ctx, eventId)
}

func (repo *Event) DeleteEventById(ctx context.Context, eventId uuid.UUID) error {
	err := repo.client.Event.
		DeleteOneID(eventId).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("event.DeleteOneID: %w", err)
	}
	return nil
}

func (repo *Event) UpdateEventById(ctx context.Context, eventId uuid.UUID, ee *ent.Event) (*ent.Event, error) {
	event, err := repo.client.Event.
		UpdateOneID(eventId).
		SetName(ee.Name).
		SetDetail(ee.Detail).
		SetLocation(ee.Location).
		SetSize(ee.Size).
		SetLimitHour(ee.LimitHour).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("event.UpdateOneID: %w", err)
	}
	return event, nil
}

func (repo *Event) GetEventAdminById(ctx context.Context, eventId uuid.UUID) (*ent.User, error) {
	event, err := repo.client.Event.
		Query().
		Where(event.ID(eventId)).
		WithAdmin().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("event.Query: %w", err)
	}
	return event.Edges.Admin, nil
}

func (repo *Event) GetEventComments(ctx context.Context, eventId uuid.UUID) ([]*ent.Comment, error) {
	comments, err := repo.client.Comment.
		Query().
		Where(comment.HasEventWith(event.ID(eventId))).
		WithUser().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("comment.Query: %w", err)
	}
	return comments, nil
}

func (repo *Event) AddNewEventParticipant(ctx context.Context, eventId uuid.UUID, eu *ent.User, comment string) error {
	user, err := repo.client.User.
		Create().
		SetName(eu.Name).
		SetIcon(eu.Icon).
		AddEventIDs(eventId).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("user.Create: %w", err)
	}
	if comment == "" {
		return nil
	}
	_, err = repo.client.Comment.
		Create().
		SetBody(comment).
		SetEventID(eventId).
		SetUserID(user.ID).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("comment.Create: %w", err)
	}
	return nil
}

func (repo *Event) ChangeEventStatusToCloseOfId(ctx context.Context, eventId uuid.UUID) (*ent.Event, error) {
	event, err := repo.client.Event.
		UpdateOneID(eventId).
		SetState(string(constant.STATE_CLOSE)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("event.UpdateOneID: %w", err)
	}
	return event, nil
}

func (repo *Event) ChangeEventStatusToCancelOfId(ctx context.Context, eventId uuid.UUID) (*ent.Event, error) {
	event, err := repo.client.Event.
		UpdateOneID(eventId).
		SetState(string(constant.STATE_CANCEL)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("event.UpdateOneID: %w", err)
	}
	return event, nil
}

func (repo *Event) GetEventUsers(ctx context.Context, eventId uuid.UUID) ([]*ent.User, error) {
	users, err := repo.client.User.
		Query().
		Where(user.HasEventsWith(event.ID(eventId))).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("user.Query: %w", err)
	}
	return users, nil
}

func (repo *Event) GetEventUsersCnt(ctx context.Context,eventId uuid.UUID) (int, error) {
	users, err := repo.client.User.
		Query().
		Where(user.HasEventsWith(event.ID(eventId))).
		All(ctx)
	if err != nil {
		return 0, fmt.Errorf("user.Query: %w", err)
	}
	return len(users),nil
}

func (repo *Event) getEventById(ctx context.Context, eventUuid uuid.UUID) (*ent.Event, error) {
	event, err := repo.client.Event.
		Query().
		Where(event.ID(eventUuid)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("event.Query: %w", err)
	}
	return event, nil
}
