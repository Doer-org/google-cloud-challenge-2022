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

type ParticipantRepository struct {
	client *ent.Client
}

func NewParticipantRepository(c *ent.Client) repository.IParticipantRepository {
	return &ParticipantRepository{
		client: c,
	}
}

func (r *ParticipantRepository) GetEventParticipants(ctx context.Context,eventId entity.EventId) ([]*entity.Participant, error) {
	eventUuid, err := uuid.Parse(string(eventId))
	if err != nil {
		return nil, fmt.Errorf("ParticipantRepository: eventUuid parse error: %w", err)
	}
	return r.getEventParticipants(ctx,eventUuid)
}

func (r *ParticipantRepository) AddNewEventParticipants(ctx context.Context, eventId entity.EventId,p *entity.Participant) ([]*entity.Participant, error) {
	eventUuid, err := uuid.Parse(string(eventId))
	if err != nil {
		return nil, fmt.Errorf("ParticipantRepository: eventUuid parse error: %w", err)
	}
	entUser, err := r.client.User.
		Create().
		SetName(p.Name).
		SetIcon(p.Icon).
		AddEventIDs(eventUuid).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("ParticipantRepository: user create query error: %w", err)
	}
	if p.Comment != "" {
		_, err = r.client.Comment.
			Create().
			SetBody(p.Comment).
			SetUserID(entUser.ID).
			SetEventID(eventUuid).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("ParticipantRepository: comment create query error: %w", err)
		}
	}
	return r.getEventParticipants(ctx,eventUuid)
}

func (r *ParticipantRepository) getEventParticipants(ctx context.Context,eventUuid uuid.UUID) ([]*entity.Participant,error) {
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
		return nil, fmt.Errorf("ParticipantRepository: get participants query error: %w", err)
	}
	entParticipantsComments, err := r.client.Comment.
		Query().
		Where(comment.HasEventWith(event.ID(eventUuid))).
		WithUser().
		WithEvent().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("ParticipantRepository: get participants comments query error: %w", err)
	}
	return EntToEntityParticipants(entParticipants,entParticipantsComments),nil
}

func EntToEntityParticipant(eu *ent.User, comment string) *entity.Participant {
	return &entity.Participant{
		Id:   entity.UserId(eu.ID.String()),
		Name: eu.Name,
		Icon: eu.Icon,
		Comment: comment,
	}
}

func EntToEntityParticipants(eus []*ent.User, ecs []*ent.Comment) []*entity.Participant {
	var ps []*entity.Participant
	for _, eu := range eus {
		hasCommentFlg := false
		for _, ec := range ecs {
			if eu.ID == ec.Edges.User.ID {
				hasCommentFlg = true
				ps = append(ps, EntToEntityParticipant(eu, ec.Body))
				break
			}
		}
		// もしコメントがなかった場合,nilを渡す
		if !hasCommentFlg {
			ps = append(ps, EntToEntityParticipant(eu, ""))
		}
	}
	return ps
}
