package persistance

import (
	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/domain/repository"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
)

type CommentRepository struct {
	client *ent.Client
}

func NewCommentRepository(c *ent.Client) repository.ICommentRepository {
	return &CommentRepository{}
}

func EntToEntityComment(e *ent.Comment) *entity.Comment {
	return &entity.Comment{
		Id:      entity.CommentId(e.ID.String()),
		UserId:  entity.UserId(e.Edges.User.ID.String()),
		EventId: entity.EventId(e.Edges.Event.ID.String()),
		Body:    e.Body,
	}
}
