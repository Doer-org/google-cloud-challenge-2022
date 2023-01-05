package handler

import (
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

func (h *UserHandler) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	j := &UserJson{}
	if err := json.NewDecoder(r.Body).Decode(j); err != nil {
		response.NewErrResponse(w, err)
		return
	}
	defer r.Body.Close()

	user, err := h.uc.CreateNewUser(
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

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	user, err := h.uc.GetUserById(r.Context(), idParam)
	if err != nil {
		response.NewErrResponse(w, err)
		return
	}
	response.ConvertToJsonResponseAndErrCheck(w, user)
}

func (h *UserHandler) GetUserByMail(w http.ResponseWriter, r *http.Request) {
	mailQuery := r.URL.Query().Get("err")
	user, err := h.uc.GetUserByMail(r.Context(), mailQuery)
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
