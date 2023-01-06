package handler

import (
	"encoding/json"
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
	j := &json_res.EventJson{}
	// TODO: bodyが空だった場合 "EOF"が入っている？
	if err := json.NewDecoder(r.Body).Decode(j); err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(err),
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
		j.Admin.Id,
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
			response.NewErrResponse(err),
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

// TODO: close,cancelのような動詞をURLに埋め込むことになるので統一すべき、
// Patchで統一
func (h *EventHandler) ChangeEventStatusToCloseOfId(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	event, err := h.uc.GetEventById(r.Context(), idParam)
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
		json_res.EntityToJsonEvent(event),
		http.StatusOK,
	)
}
