package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	json_res "github.com/Doer-org/google-cloud-challenge-2022/utils/http/json"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/http/response"
	"github.com/go-chi/chi/v5"
)

type EventHandler struct {
	eventUC usecase.IEventUsecase
	userUC  usecase.IUserUsecase
	participantUC usecase.IParticipantUsecase
}

func NewEventHandler(euc usecase.IEventUsecase,uuc usecase.IUserUsecase,puc usecase.IParticipantUsecase) *EventHandler {
	return &EventHandler{
		eventUC: euc,
		userUC: uuc,
		participantUC: puc,
	}
}

// POST /events
func (h *EventHandler) CreateNewEvent(w http.ResponseWriter, r *http.Request) {
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
	j := &json_res.EventJson{}
	if err := json.NewDecoder(r.Body).Decode(j); err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	defer r.Body.Close()

	event, err := h.eventUC.CreateNewEvent(
		r.Context(),
		j.Name,
		j.Detail,
		j.Location,
		j.Admin,
	)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(
		w,
		json_res.EntityToJsonEvent(event),
		http.StatusOK,
	)
}

// GET /events/{id}
func (h *EventHandler) GetEventById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	event, err := h.eventUC.GetEventById(r.Context(), idParam)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(
		w,
		json_res.EntityToJsonEvent(event),
		http.StatusOK,
	)
}

// DELETE /events/{id}
func (h *EventHandler) DeleteEventById(w http.ResponseWriter,r *http.Request) {
	idParam := chi.URLParam(r, "id")
	err := h.eventUC.DeleteEventById(r.Context(), idParam)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(
		w,
		fmt.Sprintf("delete event success"), //TODO: Response Objectとかあったらよさそう
		http.StatusOK,
	)
}

// PATCH /events/{id}
func (h *EventHandler) UpdateEventById(w http.ResponseWriter,r *http.Request) {
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
	j := &json_res.EventJson{}
	if err := json.NewDecoder(r.Body).Decode(j); err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	defer r.Body.Close()
	idParam := chi.URLParam(r, "id")

	event, err := h.eventUC.UpdateEventById(
		r.Context(),
		idParam,
		j.Name,
		j.Detail,
		j.Location,
	)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(
		w,
		json_res.EntityToJsonEvent(event),
		http.StatusOK,
	)
}

// GET /events/{id}/admin
func (h *EventHandler) GetEventAdminById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	admin, err := h.userUC.GetEventAdminById(r.Context(), idParam)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(
		w,
		json_res.EntityToJsonUser(admin),
		http.StatusOK,
	)
}

// GET /events/{id}/participants
func (h *EventHandler) GetEventParticipants(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	participants, err := h.participantUC.GetEventParticipants(r.Context(), idParam)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(
		w,
		json_res.EntityToJsonParticipants(participants),
		http.StatusOK,
	)
}

// POST /events/{id}/participants
func (h *EventHandler) AddNewEventParticipants(w http.ResponseWriter, r *http.Request) {
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
	j := &json_res.ParticipantJson{}
	if err := json.NewDecoder(r.Body).Decode(j); err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	defer r.Body.Close()
	idParam := chi.URLParam(r, "id")
	participants, err := h.participantUC.AddNewEventParticipants(
		r.Context(),
		idParam,
		j.Name,
		j.Comment,
	)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(
		w,
		json_res.EntityToJsonParticipants(participants),
		http.StatusOK,
	)
}

// PATCH /events/{id}/state
func (h *EventHandler) ChangeEventStatusOfId(w http.ResponseWriter, r *http.Request) {
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

	j := &json_res.EventJson{}
	if err := json.NewDecoder(r.Body).Decode(j); err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	defer r.Body.Close()
	idParam := chi.URLParam(r, "id")
	event, err := h.eventUC.ChangeEventStatusOfId(r.Context(), idParam, j.State)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	response.WriteJsonResponse(
		w,
		json_res.EntityToJsonEvent(event),
		http.StatusOK,
	)
}
