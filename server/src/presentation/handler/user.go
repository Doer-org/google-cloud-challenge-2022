package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/domain/entity"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/http/response"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	uc usecase.IUserUsecase
}

func NewUserHandler(uc usecase.IUserUsecase) *UserHandler {
	return &UserHandler{
		uc: uc,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	j := &UserJson{}
	if err := json.NewDecoder(r.Body).Decode(j); err != nil {
		response.NewErrResponse(w, err)
		return
	}
	defer r.Body.Close()

	user, err := h.uc.Create(
		r.Context(),
		j.Name,
		j.Authenticated,
		j.Mail,
		j.Icon,
	)
	if err != nil {
		response.NewErrResponse(w, err)
		return
	}
	uJson := EntityToJsonUser(user)
	response.ConvertToJsonResponseAndErrCheck(w, uJson)
}

func (h *UserHandler) GetByMail(w http.ResponseWriter, r *http.Request) {
	//TODO: getbyidとかぶる、queryとかで指定したほうがよさそう
	mailParam := chi.URLParam(r, "mail")
	user, err := h.uc.GetByMail(context.Background(), mailParam)
	if err != nil {
		response.NewErrResponse(w, err)
		return
	}
	response.ConvertToJsonResponseAndErrCheck(w, user)
}

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
