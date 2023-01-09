package repository

import (
	"context"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
)

type IParticipantRepository interface {
	GetEventParticipants(ctx context.Context,eventId entity.EventId) ([]*entity.Participant, error)
	AddNewEventParticipants(ctx context.Context, eventId entity.EventId,p *entity.Participant) ([]*entity.Participant, error)
}
