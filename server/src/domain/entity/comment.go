package entity

type CommentId string

type Comment struct {
	Id      CommentId
	UserId  UserId
	EventId EventId
	Body    string
}