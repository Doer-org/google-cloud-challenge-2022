package handler

import (
	"context"
	"encoding/json"
	"net/http"

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
	uJson := UserJson{}
	if err := json.NewDecoder(r.Body).Decode(&uJson); err != nil {
		response.NewErrResponse(w, err)
		return
	}
	defer r.Body.Close()

	err := h.uc.Create(
		r.Context(),
		uJson.Age,
		uJson.Name,
		uJson.Authenticated,
		uJson.Mail,
		uJson.Icon,
	)
	if err != nil {
		response.NewErrResponse(w, err)
		return
	}
	response.ConvertToJsonResponseAndErrCheck(
		w, response.NewResponse("user create successful"),
	)
}

func (h *UserHandler) GetByMail(w http.ResponseWriter, r *http.Request) {
	mailParam := chi.URLParam(r, "mail")
	user, err := h.uc.GetByMail(context.Background(), mailParam)
	if err != nil {
		response.NewErrResponse(w, err)
		return
	}
	response.ConvertToJsonResponseAndErrCheck(w, user)
}

type UserJson struct {
	Age           int    `json:"age"`
	Name          string `json:"name"`
	Authenticated bool   `json:"authenticated"`
	Mail          string `json:"mail"`
	Icon          string `json:"icon"`
}
