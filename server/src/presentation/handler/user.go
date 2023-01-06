package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	json_res "github.com/Doer-org/google-cloud-challenge-2022/utils/http/json"
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
	j := &json_res.UserJson{}
	if err := json.NewDecoder(r.Body).Decode(j); err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(err),
			http.StatusBadRequest,
		)
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
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(
		w,
		json_res.EntityToJsonUser(user),
		http.StatusCreated,
	)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	user, err := h.uc.GetUserById(r.Context(), idParam)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(
		w,
		json_res.EntityToJsonUser(user),
		http.StatusOK,
	)
}

func (h *UserHandler) GetUserByMail(w http.ResponseWriter, r *http.Request) {
	mailQuery := r.URL.Query().Get("mail")
	user, err := h.uc.GetUserByMail(r.Context(), mailQuery)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(
		w,
		json_res.EntityToJsonUser(user),
		http.StatusOK,
	)
}
