package usecase

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
)

type IEventUsecase interface {
	Create(ctx context.Context, name, detail, location, adminIdString string) (*entity.Event, error)
	GetById(ctx context.Context, eventIdString string) (*entity.Event, error)
}

type EventUsecase struct {
	repo repository.IEventRepository
}

func NewEventUsecae(r repository.IEventRepository) IEventUsecase {
	return &EventUsecase{
		repo: r,
	}
}

func (uc *EventUsecase) Create(ctx context.Context, name, detail, location, adminIdString string) (*entity.Event, error) {
	if name == "" {
		return nil, fmt.Errorf("EventUsecase: name is empty")
	}
	adminId := entity.UserId(adminIdString)
	if adminId == "" {
		return nil, fmt.Errorf("EventUsecase: adminId is empty")
	}
	e := &entity.Event{
		Name:     name,
		Detail:   detail,
		Location: location,
	}
	return uc.repo.Create(ctx, e, adminId)
}

func (uc *EventUsecase) GetById(ctx context.Context, eventIdString string) (*entity.Event, error) {
	// TODO: serviceとかでcastかんすを用意すべき
	eventId := entity.EventId(eventIdString)
	if eventId == "" {
		return nil, fmt.Errorf("EventUsecase: eventId parse failed")
	}
	return uc.repo.GetById(ctx, eventId)
}
