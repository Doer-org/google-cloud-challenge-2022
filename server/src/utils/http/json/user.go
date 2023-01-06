package json

import "github.com/Doer-org/google-cloud-challenge-2022/domain/entity"

type UserJson struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Authenticated bool   `json:"authenticated"`
	Mail          string `json:"mail"`
	Icon          string `json:"icon"`
}

func EntityToJsonUser(e *entity.User) *UserJson {
	return &UserJson{
		Id:            string(e.Id),
		Name:          e.Name,
		Authenticated: e.Authenticated,
		Mail:          e.Mail,
		Icon:          e.Icon,
	}
}

func EntityToJsonUsers(es []*entity.User) []*UserJson {
	var js []*UserJson
	for _, e := range es {
		js = append(js, EntityToJsonUser(e))
	}
	return js
}

type ParticipantJson struct {
	Id      string       `json:"id"`
	Name    string       `json:"name"`
	Icon    string       `json:"icon"`
	Comment *CommentJson `json:"comment"`
}

func EntityToJsonParticipant(e *entity.Participant) *ParticipantJson {
	p := &ParticipantJson{
		Id:   string(e.Id),
		Name: e.Name,
		Icon: e.Icon,
	}
	// コメントがない場合
	if e.Comment == nil {
		return p
	}
	p.Comment = EntityToJsonComment(e.Comment)
	return p
}

func EntityToJsonParticipants(es []*entity.Participant) []*ParticipantJson {
	var js []*ParticipantJson
	for _, e := range es {
		js = append(js, EntityToJsonParticipant(e))
	}
	return js
}