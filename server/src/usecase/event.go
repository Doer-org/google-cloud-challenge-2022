package usecase

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/constant"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/service"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	mycontext "github.com/Doer-org/google-cloud-challenge-2022/utils/context"
	"github.com/google/uuid"
)

type IEvent interface {
	CreateNewEvent(ctx context.Context, name, detail, location string, size, limithour int) (*ent.Event, error)
	GetEventById(ctx context.Context, eventIdString string) (*ent.Event, error)
	DeleteEventById(ctx context.Context, eventIdString string) error
	UpdateEventById(ctx context.Context, eventIdString string, name, detail, location string, size, limithour int) (*ent.Event, error)
	GetEventAdminById(ctx context.Context, eventIdString string) (*ent.User, error)
	GetEventComments(ctx context.Context, eventIdString string) ([]*ent.Comment, error)
	AddNewEventParticipant(ctx context.Context, eventIdString, name, comment string) error
	ChangeEventStatusOfId(ctx context.Context, eventIdString string, stateString string) (*ent.Event, error)
	GetEventUsers(ctx context.Context, eventIdString string) ([]*ent.User, error)
}

type Event struct {
	repo repository.IEvent
}

func NewEvent(r repository.IEvent) IEvent {
	return &Event{
		repo: r,
	}
}

func (uc *Event) CreateNewEvent(ctx context.Context, name, detail, location string, size, limithour int) (*ent.Event, error) {
	userSessId, ok := mycontext.GetUser(ctx)
	if !ok {
		return nil, fmt.Errorf("GetUser: failed to get user from context")
	}
	if name == "" {
		return nil, fmt.Errorf("name is empty")
	}
	if size == 0 {
		return nil, fmt.Errorf("size is invalid")
	}
	if limithour <= 0 || limithour >= 24 {
		return nil, fmt.Errorf("limit hour is invalid")
	}
	ee := &ent.Event{
		Name:      name,
		Detail:    detail,
		Location:  location,
		Size:      size,
		LimitHour: limithour,
	}
	return uc.repo.CreateNewEvent(ctx, userSessId, ee)
}

func (uc *Event) GetEventById(ctx context.Context, eventIdString string) (*ent.Event, error) {
	eventId, err := uuid.Parse(eventIdString)
	if err != nil {
		return nil, fmt.Errorf("eventId Parse: %w", err)
	}
	return uc.repo.GetEventById(ctx, eventId)
}

func (uc *Event) DeleteEventById(ctx context.Context, eventIdString string) error {
	eventId, err := uuid.Parse(eventIdString)
	if err != nil {
		return fmt.Errorf("eventId Parse: %w", err)
	}
	admin, err := uc.repo.GetEventAdminById(ctx, eventId)
	if err != nil {
		return fmt.Errorf("repo.GetEventAdminById: %w", err)
	}
	err = mycontext.CompareUserIdAndUserSessionId(ctx, admin.ID)
	if err != nil {
		return fmt.Errorf("compareUserIdAndUserSessionId: %w", err)
	}
	return uc.repo.DeleteEventById(ctx, eventId)
}

func (uc *Event) UpdateEventById(ctx context.Context, eventIdString string, name, detail, location string, size, limithour int) (*ent.Event, error) {
	eventId, err := uuid.Parse(eventIdString)
	if err != nil {
		return nil, fmt.Errorf("eventId Parse: %w", err)
	}
	admin, err := uc.repo.GetEventAdminById(ctx, eventId)
	if err != nil {
		return nil, fmt.Errorf("repo.GetEventAdminById: %w", err)
	}
	err = mycontext.CompareUserIdAndUserSessionId(ctx, admin.ID)
	if err != nil {
		return nil, fmt.Errorf("compareUserIdAndUserSessionId: %w", err)
	}
	// TODO: nameが空の場合は既存のnameを使う処理とかにしたほうがいいかも
	if name == "" {
		return nil, fmt.Errorf("name is empty")
	}
	if size == 0 {
		return nil, fmt.Errorf("size is invalid")
	}
	if limithour <= 0 || limithour >= 24 {
		return nil, fmt.Errorf("limit hour is invalid")
	}
	ee := &ent.Event{
		Name:     name,
		Detail:   detail,
		Location: location,
		Size:     size,
	}
	return uc.repo.UpdateEventById(ctx, eventId, ee)
}

func (uc *Event) GetEventAdminById(ctx context.Context, eventIdString string) (*ent.User, error) {
	eventId, err := uuid.Parse(eventIdString)
	if err != nil {
		return nil, fmt.Errorf("eventId Parse: %w", err)
	}
	return uc.repo.GetEventAdminById(ctx, eventId)
}

func (uc *Event) GetEventComments(ctx context.Context, eventIdString string) ([]*ent.Comment, error) {
	eventId, err := uuid.Parse(eventIdString)
	if err != nil {
		return nil, fmt.Errorf("eventId Parse: %w", err)
	}
	return uc.repo.GetEventComments(ctx, eventId)
}

func (uc *Event) AddNewEventParticipant(ctx context.Context, eventIdString, name, comment string) error {
	eventId, err := uuid.Parse(eventIdString)
	if err != nil {
		return fmt.Errorf("eventId Parse: %w", err)
	}
	event, err := uc.repo.GetEventById(ctx, eventId)
	if err != nil {
		return fmt.Errorf("repo.GetEventById: %w", err)
	}
	if event.State != constant.STATE_OPEN {
		return fmt.Errorf("state is not open")
	}
	nowSize, err := uc.repo.GetEventUsersCnt(ctx, eventId)
	if err != nil {
		return fmt.Errorf("repo.GetEventUsersCnt: %w", err)
	}
	if event.Size <= nowSize {
		return fmt.Errorf("the room is already full")
	}
	if name == "" {
		return fmt.Errorf("name is empty")
	}
	eu := &ent.User{
		Name: name,
		Icon: service.GetRandomDefaultIcon(),
	}
	err = uc.repo.AddNewEventParticipant(ctx, eventId, eu, comment)
	if err != nil {
		return fmt.Errorf("addNewEventParticipant: %w", err)
	}
	return nil
}

func (uc *Event) ChangeEventStatusOfId(ctx context.Context, eventIdString string, state string) (*ent.Event, error) {
	eventId, err := uuid.Parse(eventIdString)
	if err != nil {
		return nil, fmt.Errorf("eventId Parse: %w", err)
	}
	admin, err := uc.repo.GetEventAdminById(ctx, eventId)
	if err != nil {
		return nil, fmt.Errorf("repo.GetEventAdminById: %w", err)
	}
	err = mycontext.CompareUserIdAndUserSessionId(ctx, admin.ID)
	if err != nil {
		return nil, fmt.Errorf("compareUserIdAndUserSessionId: %w", err)
	}
	event, err := uc.repo.GetEventById(ctx, eventId)
	if err != nil {
		return nil, fmt.Errorf("repo.GetEventById: %w", err)
	}
	if event.State != constant.STATE_OPEN {
		return nil, fmt.Errorf("state is not open")
	}
	if state == "" {
		return nil, fmt.Errorf("state is Empty")
	}
	switch state {
	case constant.STATE_CLOSE:
		return uc.repo.ChangeEventStatusToCloseOfId(ctx, eventId)
	case constant.STATE_CANCEL:
		return uc.repo.ChangeEventStatusToCancelOfId(ctx, eventId)
	default:
		return nil, fmt.Errorf("received state is not matched")
	}
}

func (uc *Event) GetEventUsers(ctx context.Context, eventIdString string) ([]*ent.User, error) {
	eventId, err := uuid.Parse(eventIdString)
	if err != nil {
		return nil, fmt.Errorf("eventId Parse: %w", err)
	}
	return uc.repo.GetEventUsers(ctx, eventId)
}
