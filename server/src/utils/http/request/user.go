package request

type UserRequestJson struct {
	Name          string `json:"name"`
	Authenticated bool   `json:"authenticated"`
	Mail          string `json:"mail"`
	Icon          string `json:"icon"`
}
