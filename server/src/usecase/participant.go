package usecase

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
)

type IParticipantUsecase interface{
	CreateNewParticipant(ctx context.Context, name string, body string, eventIdString string) (*entity.Participant, error)
}

type ParticipantUsecase struct {
	repo repository.IParticipantRepository
}

func NewParticipantUsecase(repo repository.IParticipantRepository) IParticipantUsecase {
	return &ParticipantUsecase{
		repo: repo,
	}
}

func (uc *ParticipantUsecase) CreateNewParticipant(ctx context.Context, name string, body string, eventIdString string) (*entity.Participant, error) {
	eventId := entity.EventId(eventIdString)
	if eventId == "" {
		return nil, fmt.Errorf("ParticipantUsecase: eventId parse failed")
	}

	// iconはランダム
	icon := "random"

	if name == "" {
		return nil, fmt.Errorf("ParticipantUsecase: name is empty")
	}
	p := &entity.Participant{
		Name: name,
		Icon: icon,
		Comment: &entity.Comment{
			Body: body,
		},
	}
	return uc.repo.CreateNewParticipant(ctx,p,eventId)
}
