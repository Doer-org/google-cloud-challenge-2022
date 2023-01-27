package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	res "github.com/Doer-org/google-cloud-challenge-2022/presentation/http/response"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	"github.com/go-chi/chi/v5"
)

type Event struct {
	uc usecase.IEvent
}

func NewEvent(uc usecase.IEvent) *Event {
	return &Event{
		uc: uc,
	}
}

// POST /events
func (h *Event) CreateNewEvent(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: request body is empty")), http.StatusBadRequest)
		return
	}
	var j eventJson
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: Decoder: %w", err)), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	event, err := h.uc.CreateNewEvent(
		r.Context(),
		j.Name,
		j.Detail,
		j.Location,
		j.Size,
		j.LimitTime,
	)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: CreateNewEvent: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, event, http.StatusOK)
}

// GET /events/{id}
func (h *Event) GetEventById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	event, err := h.uc.GetEventById(r.Context(), idParam)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetEventById: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, event, http.StatusOK)
}

// DELETE /events/{id}
func (h *Event) DeleteEventById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	err := h.uc.DeleteEventById(r.Context(), idParam)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: DeleteEventById: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, res.New200SuccessJson("success: delete event"), http.StatusOK)
}

// PATCH /events/{id}
func (h *Event) UpdateEventById(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: request body is empty")), http.StatusBadRequest)
		return
	}
	var j eventJson
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
		j.Size,
		j.LimitTime,
	)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: UpdateEventById: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, event, http.StatusNoContent)
}

// GET /events/{id}/admin
func (h *Event) GetEventAdminById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	admin, err := h.uc.GetEventAdminById(r.Context(), idParam)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetEventAdminById: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, admin, http.StatusOK)
}

// GET /events/{id}/comments
func (h *Event) GetEventComments(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	comments, err := h.uc.GetEventComments(r.Context(), idParam)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetEventComments: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, comments, http.StatusOK)
}

// PATCH /events/{id}/state
func (h *Event) ChangeEventStatusOfId(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: request body is empty")), http.StatusBadRequest)
		return
	}
	// TODO: entに変えてもよさそう
	var j eventJson
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
func (h *Event) AddNewEventParticipant(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: request body is empty")), http.StatusBadRequest)
		return
	}
	var j participantJson
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
func (h *Event) GetEventUsers(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	users, err := h.uc.GetEventUsers(r.Context(), idParam)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetEventUsers: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, users, http.StatusOK)
}

type eventJson struct {
	Name      string `json:"name"`
	Detail    string `json:"detail"`
	Location  string `json:"location"`
	Size      int    `json:"size"`
	LimitTime time.Time    `json:"limit_time"`
	State     string `json:"state"`
	Type      string `json:"type"`
}

type participantJson struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}
