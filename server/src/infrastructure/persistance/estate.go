package persistance

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
)

type EStateRepository struct {
	client *ent.Client
}

func NewEStateRepository(c *ent.Client) repository.IEStateRepository {
	return &EStateRepository{
		client: c,
	}
}

func (r *EStateRepository) Create(ctx context.Context) (*entity.EState, error) {
	entEState, err := r.client.EState.
		Create().
		SetName("open").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("EStateRepository: create estate query error: %w", err)
	}
	return EntToEntityEState(entEState), nil
}

func (r *EStateRepository) UpdateStatusClose(ctx context.Context, id entity.EStateId) (*entity.EState, error) {
	uuid, err := uuid.Parse(string(id))
	if err != nil {
		return nil, fmt.Errorf("EStateRepository: uuid parse error: %w", err)
	}
	entEState, err := r.client.EState.
		UpdateOneID(uuid).
		SetName("close").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("EStateRepository: update estate query error: %w", err)
	}
	return EntToEntityEState(entEState), nil
}

func (r *EStateRepository) UpdateStatusCancel(ctx context.Context, id entity.EStateId) (*entity.EState, error) {
	uuid, err := uuid.Parse(string(id))
	if err != nil {
		return nil, fmt.Errorf("EStateRepository: uuid parse error: %w", err)
	}
	entEState, err := r.client.EState.
		UpdateOneID(uuid).
		SetName("cancel").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("EStateRepository: update estate query error: %w", err)
	}
	return EntToEntityEState(entEState), nil
}

func EntToEntityEState(e *ent.EState) *entity.EState {
	return &entity.EState{
		Id:   entity.EStateId(e.ID.String()),
		Name: e.Name,
	}
}

func EntityToEntEState(es *entity.EState) *ent.EState {
	return &ent.EState{
		Name: es.Name,
	}
}
