package json

import "github.com/Doer-org/google-cloud-challenge-2022/domain/entity"

type ParticipantJson struct {
	Id      string       `json:"id"`
	Name    string       `json:"name"`
	Icon    string       `json:"icon"`
	Comment string       `json:"comment"`
}

func EntityToJsonParticipant(e *entity.Participant) *ParticipantJson {
	return &ParticipantJson{
		Id:   string(e.Id),
		Name: e.Name,
		Icon: e.Icon,
		Comment: e.Comment,
	}
}

func EntityToJsonParticipants(es []*entity.Participant) []*ParticipantJson {
	var js []*ParticipantJson
	for _, e := range es {
		js = append(js, EntityToJsonParticipant(e))
	}
	return js
}
