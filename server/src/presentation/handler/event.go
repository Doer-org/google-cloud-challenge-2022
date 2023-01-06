package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
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
	eJson := &EventJson{}
	if err := json.NewDecoder(r.Body).Decode(eJson); err != nil {
		response.NewErrResponse(w, err)
		return
	}
	defer r.Body.Close()

	event, err := h.uc.CreateNewEvent(
		r.Context(),
		eJson.Name,
		eJson.Detail,
		eJson.Location,
		eJson.AdminId,
	)
	if err != nil {
		response.NewErrResponse(w, err)
		return
	}
	response.ConvertToJsonResponseAndErrCheck(w, event)
}

func (h *EventHandler) GetEventById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	event, err := h.uc.GetEventById(r.Context(), idParam)
	if err != nil {
		response.NewErrResponse(w, err)
		return
	}
	response.ConvertToJsonResponseAndErrCheck(w, event)
}

// TODO: close,cancelのような動詞をURLに埋め込むことになるので統一すべき、
// Patchで統一
func (h *EventHandler) ChangeEventStatusToCloseOfId(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	event, err := h.uc.GetEventById(r.Context(), idParam)
	if err != nil {
		response.NewErrResponse(w, err)
		return
	}
	response.ConvertToJsonResponseAndErrCheck(w, event)
}

type EventJson struct {
	Id       string        `json:"id"`
	Name     string        `json:"name"`
	Detail   string        `json:"detail"`
	Location string        `json:"location"`
	AdminId  string        `json:"admin_id"`
	State    string        `json:"state"`
	Type     string        `json:"type"`
	Comments []CommentJson `json:"comments"`
}
