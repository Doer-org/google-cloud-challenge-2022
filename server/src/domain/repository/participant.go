package repository

import (
	"context"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
)

type IParticipantRepository interface {
	CreateNewParticipant(ctx context.Context, p *entity.Participant, eventId entity.EventId) (*entity.Participant, error)
}
