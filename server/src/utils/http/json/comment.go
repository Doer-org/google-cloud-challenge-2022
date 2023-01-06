package json

import "github.com/Doer-org/google-cloud-challenge-2022/domain/entity"

type CommentJson struct {
	Id   string `json:"id"`
	Body string `json:"body"`
}

func EntityToJsonComment(e *entity.Comment) *CommentJson {
	return &CommentJson{
		Id:   string(e.Id),
		Body: e.Body,
	}
}
