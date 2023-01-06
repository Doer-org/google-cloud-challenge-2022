package json

type EventJson struct {
	Id       string        `json:"id"`
	Name     string        `json:"name"`
	Detail   string        `json:"detail"`
	Location string        `json:"location"`
	AdminId  string        `json:"admin_id"`
	State    string        `json:"state"`
	Type     string        `json:"type"`
	Comments []CommentJson `json:"comments"`
}

