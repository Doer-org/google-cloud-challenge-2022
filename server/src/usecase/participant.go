package usecase

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/service"
)

type IParticipantUsecase interface {
	GetEventParticipants(ctx context.Context, eventIdString string) ([]*entity.Participant, error)
	AddNewEventParticipants(ctx context.Context, eventIdString ,name,comment string) ([]*entity.Participant, error)
}

type ParticipantUsecase struct {
	repo repository.IParticipantRepository
}

func NewParticipantUsecase(repo repository.IParticipantRepository) IParticipantUsecase {
	return &ParticipantUsecase{
		repo: repo,
	}
}

func (uc *ParticipantUsecase) GetEventParticipants(ctx context.Context, eventIdString string) ([]*entity.Participant, error) {
	eventId := entity.EventId(eventIdString)
	if eventId == "" {
		return nil, fmt.Errorf("ParticipantUsecase: eventId parse failed")
	}
	return uc.repo.GetEventParticipants(ctx,eventId)
}

func (uc *ParticipantUsecase) AddNewEventParticipants(ctx context.Context, eventIdString ,name,comment string) ([]*entity.Participant, error) {
	eventId := entity.EventId(eventIdString)
	if eventId == "" {
		return nil, fmt.Errorf("ParticipantUsecase: eventId parse failed")
	}
	if name == "" {
		return nil, fmt.Errorf("ParticipantUsecase: name is empty")
	}
	p := &entity.Participant{
		Name: name,
		Icon: service.GetRandomDefaultIcon(),
		Comment: comment,
	}
	return uc.repo.AddNewEventParticipants(ctx,eventId,p)
}
