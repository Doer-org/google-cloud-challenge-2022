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

type EventHandler struct {
	uc usecase.IEventUsecase
}

// TODO: handlerとかusecaseつけなくてもいい
// TODO: domainをけす?
func NewEventHandler(uc usecase.IEventUsecase) *EventHandler {
	return &EventHandler{
		uc: uc,
	}
}

// POST /events
func (h *EventHandler) CreateNewEvent(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: request body is empty")), http.StatusBadRequest)
		return
	}
	var j request.EventJson
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: Decoder: %w", err)), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	event, err := h.uc.CreateNewEvent(
		r.Context(),
		j.Admin,
		j.Name,
		j.Detail,
		j.Location,
	)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: CreateNewEvent: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, event, http.StatusOK)
}

// GET /events/{id}
func (h *EventHandler) GetEventById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	event, err := h.uc.GetEventById(r.Context(), idParam)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetEventById: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, event, http.StatusOK)
}

// DELETE /events/{id}
func (h *EventHandler) DeleteEventById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	err := h.uc.DeleteEventById(r.Context(), idParam)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: DeleteEventById: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, res.New200SuccessJson("success: delete event"), http.StatusOK)
}

// PATCH /events/{id}
func (h *EventHandler) UpdateEventById(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: request body is empty")), http.StatusBadRequest)
		return
	}
	var j request.EventJson
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: Decoder: %w", err)), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	idParam := chi.URLParam(r, "id")
	event, err := h.uc.UpdateEventById(
		r.Context(),
		idParam,
		j.Name,
		j.Detail,
		j.Location,
	)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: UpdateEventById: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, event, http.StatusOK)
}

// GET /events/{id}/admin
func (h *EventHandler) GetEventAdminById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	admin, err := h.uc.GetEventAdminById(r.Context(), idParam)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetEventAdminById: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, admin, http.StatusOK)
}

// GET /events/{id}/comments
func (h *EventHandler) GetEventComments(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	comments, err := h.uc.GetEventComments(r.Context(), idParam)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetEventComments: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, comments, http.StatusOK)
}

// PATCH /events/{id}/state
func (h *EventHandler) ChangeEventStatusOfId(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: request body is empty")), http.StatusBadRequest)
		return
	}
	var j request.EventJson
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: Decoder: %w", err)), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	idParam := chi.URLParam(r, "id")
	event, err := h.uc.ChangeEventStatusOfId(r.Context(), idParam, j.State)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: ChangeEventStatusOfId: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, event, http.StatusOK)
}

// POST /events/{id}/participants
func (h *EventHandler) AddNewEventParticipant(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: request body is empty")), http.StatusBadRequest)
		return
	}
	var j request.ParticipantJson
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: Decode: %w", err)), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	idParam := chi.URLParam(r, "id")
	err := h.uc.AddNewEventParticipant(
		r.Context(),
		idParam,
		j.Name,
		j.Comment,
	)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: AddNewEventParticipant: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, res.New200SuccessJson("success: add participant"), http.StatusOK)
}

// GET /events/{id}/users
func (h *EventHandler) GetEventUsers(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	users, err := h.uc.GetEventUsers(r.Context(), idParam)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetEventUsers: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, users, http.StatusOK)
}
