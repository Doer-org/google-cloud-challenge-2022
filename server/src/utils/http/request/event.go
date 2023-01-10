package request

type EventRequestJson struct {
	Name     string `json:"name"`
	Detail   string `json:"detail"`
	Location string `json:"location"`
	Admin    string `json:"admin"`
	State    string `json:"state"`
	Type     string `json:"type"`
}
