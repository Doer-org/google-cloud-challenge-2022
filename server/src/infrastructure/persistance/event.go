package persistance

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent/ecomment"
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

func (r *EventRepository) Create(ctx context.Context, e *entity.Event, adminId entity.UserId) (*entity.Event, error) {
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
		AddUserIDs(adminUuid).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: create event query error: %w", err)
	}
	_, err = r.client.EState.
		Create().
		SetName("open").
		SetEvent(ee).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: create estate query error: %w", err)
	}
	_, err = r.client.EType.
		Create().
		SetName("once").
		SetEvent(ee).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: create etype query error: %w", err)
	}
	entEvent, err := r.client.Event.
		Query().
		Where(event.ID(ee.ID)).
		WithAdmin().
		WithType().
		WithState().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: get event query error: %w", err)
	}
	// イベント作成のタイミングでコメントが入っていることはないので空のオブジェクトを渡している
	return EntToEntityEvent(entEvent, []*ent.User{entEvent.Edges.Admin}, []*ent.Ecomment{}), nil
}

func (r *EventRepository) GetById(ctx context.Context, id entity.EventId) (*entity.Event, error) {
	eventUuid, err := uuid.Parse(string(id))
	if err != nil {
		return nil, fmt.Errorf("EventRepository: uuid parse error: %w", err)
	}
	entEvent, err := r.client.Event.
		Query().
		Where(event.ID(eventUuid)).
		WithState().
		WithType().
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
	entParticipantsComments, err := r.client.Ecomment.
		Query().
		Where(ecomment.HasEventWith(event.ID(eventUuid))).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("EventRepository: get participants comments query error: %w", err)
	}
	return EntToEntityEvent(entEvent, entParticipants, entParticipantsComments), nil
}

func EntToEntityEvent(ee *ent.Event, eus []*ent.User, ecs []*ent.Ecomment) *entity.Event {
	return &entity.Event{
		Id:           entity.EventId(ee.ID.String()),
		Name:         ee.Name,
		Detail:       ee.Detail,
		Location:     ee.Location,
		Admin:        EntToEntityUser(ee.Edges.Admin),
		State:        EntToEntityEState(ee.Edges.State),
		Type:         EntToEntityEType(ee.Edges.Type),
		Participants: EntsToEntitiesParticipant(eus, ecs),
	}
}
