package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/request"
	res "github.com/Doer-org/google-cloud-challenge-2022/presentation/http/response"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	"github.com/go-chi/chi/v5"
)

type User struct {
	uc usecase.IUser
}

func NewUser(uc usecase.IUser) *User {
	return &User{
		uc: uc,
	}
}

// POST /users
func (h *User) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: request body is empty")), http.StatusBadRequest)
		return
	}
	var j request.UserJson
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: Decode: %w", err)), http.StatusBadRequest)
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
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: CreateNewUser: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, user, http.StatusCreated)
}

// GET /users/{id}
func (h *User) GetUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	user, err := h.uc.GetUserById(r.Context(), idParam)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetUserById: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, user, http.StatusOK)
}

// DELETE /users/{id}
func (h *User) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	err := h.uc.DeleteUserById(r.Context(), idParam)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: DeleteUserById: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, res.New200SuccessJson("delete user success"), http.StatusOK)
}

// PATCH /users/{id}
func (h *User) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: request body is empty")), http.StatusBadRequest)
		return
	}
	var j request.UserJson
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: Decode: %w", err)), http.StatusBadRequest)
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
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: UpdateUserById: %w", err)), http.StatusBadRequest)
		return
	}
	//TODO: udpateはcreated?
	res.WriteJson(w, user, http.StatusCreated)
}

// GET /users/{id}/events
func (h *User) GetUserEvents(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	events, err := h.uc.GetUserEvents(r.Context(), idParam)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetUserEvents: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, events, http.StatusOK)
}

// TODO: openapiに追加する
// GET /users?mail=<user mail>
func (h *User) GetUserByMail(w http.ResponseWriter, r *http.Request) {
	mailQuery := r.URL.Query().Get("mail")
	user, err := h.uc.GetUserByMail(r.Context(), mailQuery)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetUserByMail: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, user, http.StatusOK)
}
