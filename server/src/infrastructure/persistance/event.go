package persistance

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
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

func (r *EventRepository) CreateNewEvent(ctx context.Context, e *entity.Event, adminId entity.UserId) (*entity.Event, error) {
	adminUuid, err := uuid.Parse(string(adminId))
	if err != nil {
		return nil, fmt.Errorf("EventRepository: adminUuid parse error: %w", err)
	}
	ee, err := r.client.Event.
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
	event, err := r.getEventById(ctx, ee.ID)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: getEventById error: %w", err)
	}
	return event, nil
}

func (r *EventRepository) GetEventById(ctx context.Context, eventId entity.EventId) (*entity.Event, error) {
	eventUuid, err := uuid.Parse(string(eventId))
	if err != nil {
		return nil, fmt.Errorf("EStateRepository: uuid parse error: %w", err)
	}
	event, err := r.getEventById(ctx, eventUuid)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: getEventById error: %w", err)
	}
	return event, nil
}

func (r *EventRepository) ChangeEventStatusToCloseOfId(ctx context.Context, eventId entity.EventId) (*entity.Event, error) {
	eventUuid, err := uuid.Parse(string(eventId))
	if err != nil {
		return nil, fmt.Errorf("EStateRepository: uuid parse error: %w", err)
	}
	_, err = r.client.Event.
		UpdateOneID(eventUuid).
		SetState(string(entity.CLOSE_STATE)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: update event query error: %w", err)
	}
	event, err := r.getEventById(ctx, eventUuid)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: getEventById error: %w", err)
	}
	return event, nil
}

func (r *EventRepository) ChangeEventStatusToCancelOfId(ctx context.Context, eventId entity.EventId) (*entity.Event, error) {
	eventUuid, err := uuid.Parse(string(eventId))
	if err != nil {
		return nil, fmt.Errorf("EStateRepository: uuid parse error: %w", err)
	}
	_, err = r.client.Event.
		UpdateOneID(eventUuid).
		SetState(string(entity.CANCEL_STATE)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: update event query error: %w", err)
	}
	event, err := r.getEventById(ctx, eventUuid)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: getEventById error: %w", err)
	}
	return event, nil
}

func (r *EventRepository) getEventById(ctx context.Context, eventUuid uuid.UUID) (*entity.Event, error) {
	entEvent, err := r.client.Event.
		Query().
		Where(event.ID(eventUuid)).
		WithAdmin().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: get event query error: %w", err)
	}
	// 指定したeventUuidのイベントに参加しているユーザーをすべて取得する
	entParticipants, err := r.client.User.
		Query().
		Where(func(s *sql.Selector) {
			t := sql.Table(user.EventsTable)
			s.LeftJoin(t).On(s.C(user.FieldID), t.C(user.EventsPrimaryKey[1]))
		}).
		Where(user.HasEventsWith(event.ID(eventUuid))).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: get participants query error: %w", err)
	}
	entParticipantsComments, err := r.client.Comment.
		Query().
		Where(comment.HasEventWith(event.ID(eventUuid))).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: get participants comments query error: %w", err)
	}
	return EntToEntityEvent(entEvent, entParticipants, entParticipantsComments), nil
}

func EntToEntityEvent(ee *ent.Event, eus []*ent.User, ecs []*ent.Comment) *entity.Event {
	return &entity.Event{
		Id:           entity.EventId(ee.ID.String()),
		Name:         ee.Name,
		Detail:       ee.Detail,
		Location:     ee.Location,
		Admin:        EntToEntityUser(ee.Edges.Admin),
		State:        ee.State,
		Type:         ee.Type,
		Participants: EntsToEntitiesParticipant(eus, ecs),
	}
}
