package entity

type ECommentId string

// TODO: ent配下のスキーマがEcommentになっている
// ECommentにしたい
type EComment struct {
	Id      ECommentId
	UserId  UserId
	EventId EventId
	Body    string
}
