package persistance

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
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

func (r *ParticipantRepository) CreateNewParticipant(ctx context.Context, p *entity.Participant, eventId entity.EventId) (*entity.Participant, error) {
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
	if p.Comment == nil {
		return EntToEntityParticipant(entUser, nil), nil
	}
	entComment, err := r.client.Comment.
		Create().
		SetBody(p.Comment.Body).
		SetUserID(entUser.ID).
		SetEventID(eventUuid).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("ParticipantRepository: comment create query error: %w", err)
	}
	return EntToEntityParticipant(entUser, entComment), nil
}

func EntToEntityParticipant(eu *ent.User, ec *ent.Comment) *entity.Participant {
	p := &entity.Participant{
		Id:   entity.UserId(eu.ID.String()),
		Name: eu.Name,
		Icon: eu.Icon,
	}
	//コメントがない場合
	if ec == nil {
		return p
	}
	c := EntToEntityComment(ec)
	p.Comment = c
	return p
}

// TODO entの設計がいけてないかもしれない
// 本来はeu.Edges.Commentでアクセスできるべき...?
func EntToEntityParticipants(eus []*ent.User, ecs []*ent.Comment) []*entity.Participant {
	var ps []*entity.Participant
	for _, eu := range eus {
		hasCommentFlg := false
		for _, ec := range ecs {
			if eu.ID == ec.Edges.User.ID {
				hasCommentFlg = true
				ps = append(ps, EntToEntityParticipant(eu, ec))
				break
			}
		}
		// もしコメントがなかった場合,nilを渡す
		if !hasCommentFlg {
			ps = append(ps, EntToEntityParticipant(eu, nil))
		}
	}
	return ps
}
