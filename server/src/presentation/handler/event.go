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
	uc usecase.IEventUsecase
}

func NewEventHandler(uc usecase.IEventUsecase) *EventHandler {
	return &EventHandler{
		uc: uc,
	}
}

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

	event, err := h.uc.CreateNewEvent(
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

func (h *EventHandler) GetEventById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	event, err := h.uc.GetEventById(r.Context(), idParam)
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
	event, err := h.uc.ChangeEventStatusOfId(r.Context(), idParam, j.State)
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
