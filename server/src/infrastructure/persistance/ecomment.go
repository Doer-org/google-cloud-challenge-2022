package persistance

import (
	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
)

type ECommentRepository struct {
	client *ent.Client
}

func NewECommentRepository(c *ent.Client) repository.IECommentRepository {
	return &ECommentRepository{}
}

func EntToEntityEComment(e *ent.Ecomment) *entity.EComment {
	return &entity.EComment{
		Id:      entity.ECommentId(e.ID.String()),
		UserId:  entity.UserId(e.Edges.User.ID.String()),
		EventId: entity.EventId(e.Edges.Event.ID.String()),
		Body:    e.Body,
	}
}
