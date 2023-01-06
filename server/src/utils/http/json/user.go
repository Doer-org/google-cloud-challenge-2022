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
