package handler

type CommentJson struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	Body   string `json:"body"`
}
