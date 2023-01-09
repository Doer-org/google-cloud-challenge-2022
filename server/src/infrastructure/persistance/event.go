package persistance

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
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

func (r *EventRepository) CreateNewEvent(ctx context.Context, e *entity.Event, adminId entity.UserId) (*entity.Event, error) {
	adminUuid, err := uuid.Parse(string(adminId))
	if err != nil {
		return nil, fmt.Errorf("EventRepository: adminUuid parse error: %w", err)
	}
	entEvent, err := r.client.Event.
		Create().
		SetName(e.Name).
		SetDetail(e.Detail).
		SetLocation(e.Location).
		SetAdminID(adminUuid).
		SetType(string(entity.ONCE_TYPE)).
		SetState(string(entity.OPEN_STATE)).
		AddUserIDs(adminUuid).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: create event query error: %w", err)
	}
	return r.getEventById(ctx, entEvent.ID)
}

func (r *EventRepository) GetEventById(ctx context.Context, eventId entity.EventId) (*entity.Event, error) {
	eventUuid, err := uuid.Parse(string(eventId))
	if err != nil {
		return nil, fmt.Errorf("EventRepository: uuid parse error: %w", err)
	}
	return r.getEventById(ctx, eventUuid)
}

func (r *EventRepository) DeleteEventById(ctx context.Context, eventId entity.EventId) error {
	eventUuid, err := uuid.Parse(string(eventId))
	if err != nil {
		return fmt.Errorf("EventRepository: uuid parse error: %w", err)
	}
	err = r.client.Event.
		DeleteOneID(eventUuid).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("EventRepository: delete event query error: %w", err)
	}
	return nil
}

func (r *EventRepository) UpdateEventById(ctx context.Context, eventId entity.EventId, e *entity.Event) (*entity.Event, error) {
	eventUuid, err := uuid.Parse(string(eventId))
	if err != nil {
		return nil,fmt.Errorf("EventRepository: uuid parse error: %w", err)
	}
	entEvent,err := r.client.Event.
		UpdateOneID(eventUuid).
		SetName(e.Name).
		SetDetail(e.Detail).
		SetLocation(e.Location).
		Save(ctx)
	if err != nil {
		return nil,fmt.Errorf("EventRepository: update event query error: %w", err)
	}
	return EntToEntityEvent(entEvent),nil
}

func (r *EventRepository) ChangeEventStatusToCloseOfId(ctx context.Context, eventId entity.EventId) (*entity.Event, error) {
	eventUuid, err := uuid.Parse(string(eventId))
	if err != nil {
		return nil, fmt.Errorf("EventRepository: uuid parse error: %w", err)
	}
	_, err = r.client.Event.
		UpdateOneID(eventUuid).
		SetState(string(entity.CLOSE_STATE)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: update event query error: %w", err)
	}
	return r.getEventById(ctx, eventUuid)
}

func (r *EventRepository) ChangeEventStatusToCancelOfId(ctx context.Context, eventId entity.EventId) (*entity.Event, error) {
	eventUuid, err := uuid.Parse(string(eventId))
	if err != nil {
		return nil, fmt.Errorf("EventRepository: uuid parse error: %w", err)
	}
	_, err = r.client.Event.
		UpdateOneID(eventUuid).
		SetState(string(entity.CANCEL_STATE)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: update event query error: %w", err)
	}
	return r.getEventById(ctx, eventUuid)
}

func (r *EventRepository) GetUserEvents(ctx context.Context, userId entity.UserId) ([]*entity.Event,error){
	userUuid, err := uuid.Parse(string(userId))
	if err != nil {
		return nil,fmt.Errorf("EventRepository: userUuid parse error: %w", err)
	}
	entUser,err := r.client.User.
		Query().
		Where(user.ID(userUuid)).
		WithEvents().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: get user event query error: %w", err)
	}
	return EntToEntityEvents(entUser.Edges.Events),nil
}

func (r *EventRepository) getEventById(ctx context.Context, eventUuid uuid.UUID) (*entity.Event, error) {
	entEvent, err := r.client.Event.
		Query().
		Where(event.ID(eventUuid)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: get event query error: %w", err)
	}
	return EntToEntityEvent(entEvent), nil
}

func EntToEntityEvent(ee *ent.Event) *entity.Event {
	return &entity.Event{
		Id:           entity.EventId(ee.ID.String()),
		Name:         ee.Name,
		Detail:       ee.Detail,
		Location:     ee.Location,
		State:        ee.State,
		Type:         ee.Type,
	}
}

func EntToEntityEvents(ees []*ent.Event) []*entity.Event {
	var es []*entity.Event
	for _,ee := range ees {
		es = append(es, EntToEntityEvent(ee))
	}
	return es
}
