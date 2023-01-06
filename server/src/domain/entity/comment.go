package entity

type CommentId string

type Comment struct {
	Id      CommentId
	Body    string
}
