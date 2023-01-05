package persistance

import (
	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
)

type ETypeRepository struct {
	client *ent.Client
}

func NewETypeRepository(c *ent.Client) repository.IETypeRepository {
	return &ETypeRepository{}
}

func EntToEntityEType(e *ent.EType) *entity.EType {
	return &entity.EType{
		Id:   entity.ETypeId(e.ID.String()),
		Name: e.Name,
	}
}
