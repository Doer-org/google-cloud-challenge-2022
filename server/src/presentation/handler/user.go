package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/http/request"
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

// POST /users
func (h *UserHandler) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(
				http.StatusBadRequest,
				"StatusBadRequest",
				fmt.Errorf("request body is empty"),
			),
			http.StatusBadRequest,
		)
		return
	}
	var j request.UserRequestJson
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest, "StatusBadRequest", err),
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
			response.NewErrResponse(http.StatusBadRequest, "StatusBadRequest", err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(w, user, http.StatusCreated)
}

// GET /users/{id}
func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	user, err := h.uc.GetUserById(r.Context(), idParam)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest, "StatusBadRequest", err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(w, user, http.StatusOK)
}

// DELETE /users/{id}
func (h *UserHandler) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	err := h.uc.DeleteUserById(r.Context(), idParam)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest, "StatusBadRequest", err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(w, fmt.Sprintf("delete user success"), http.StatusOK)
}

// PATCH /users/{id}
func (h *UserHandler) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(
				http.StatusBadRequest,
				"StatusBadRequest",
				fmt.Errorf("request body is empty"),
			),
			http.StatusBadRequest,
		)
		return
	}
	var j request.UserRequestJson
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest, "StatusBadRequest", err),
			http.StatusBadRequest,
		)
		return
	}
	defer r.Body.Close()
	idParam := chi.URLParam(r, "id")
	user, err := h.uc.UpdateUserById(
		r.Context(),
		idParam,
		j.Name,
		j.Authenticated,
		j.Mail,
		j.Icon,
	)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest, "StatusBadRequest", err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(w, user, http.StatusCreated)
}

// GET /users/{id}/events
func (h *UserHandler) GetUserEvents(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	events, err := h.uc.GetUserEvents(r.Context(), idParam)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest, "StatusBadRequest", err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(w, events, http.StatusOK)
}

// TODO: openapiに追加する
// GET /users?mail=<user mail>
func (h *UserHandler) GetUserByMail(w http.ResponseWriter, r *http.Request) {
	mailQuery := r.URL.Query().Get("mail")
	user, err := h.uc.GetUserByMail(r.Context(), mailQuery)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest, "StatusBadRequest", err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(w, user, http.StatusOK)
}
