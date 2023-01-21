package request

type EventJson struct {
	Name     string `json:"name"`
	Detail   string `json:"detail"`
	Location string `json:"location"`
	Size     int    `json:"size"`
	State    string `json:"state"`
	Type     string `json:"type"`
}
