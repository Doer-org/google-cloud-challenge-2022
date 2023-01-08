package json

import "github.com/Doer-org/google-cloud-challenge-2022/domain/entity"

type EventJson struct {
	Id           string             `json:"id"`
	Name         string             `json:"name"`
	Detail       string             `json:"detail"`
	Location     string             `json:"location"`
	Admin        string             `json:"admin"`
	State        string             `json:"state"`
	Type         string             `json:"type"`
}

func EntityToJsonEvent(e *entity.Event) *EventJson {
	return &EventJson{
		Id:           string(e.Id),
		Name:         e.Name,
		Detail:       e.Detail,
		Location:     e.Location,
		State:        e.State,
		Type:         e.Type,
	}
}

func EntityToJsonEvents(es []*entity.Event) []*EventJson {
	var js []*EventJson
	for _,e := range es {
		js = append(js, EntityToJsonEvent(e))
	}
	return js
}
