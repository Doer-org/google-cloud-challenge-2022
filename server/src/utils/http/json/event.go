package json

import "github.com/Doer-org/google-cloud-challenge-2022/domain/entity"

type EventJson struct {
	Id           string             `json:"id"`
	Name         string             `json:"name"`
	Detail       string             `json:"detail"`
	Location     string             `json:"location"`
	Admin        *UserJson          `json:"admin"`
	State        string             `json:"state"`
	Type         string             `json:"type"`
	Participants []*ParticipantJson `json:"participants"`
}

func EntityToJsonEvent(e *entity.Event) *EventJson {
	return &EventJson{
		Id:           string(e.Id),
		Name:         e.Name,
		Detail:       e.Detail,
		Location:     e.Location,
		Admin:        EntityToJsonUser(e.Admin),
		State:        e.State,
		Type:         e.Type,
		Participants: EntityToJsonParticipants(e.Participants),
	}
}
