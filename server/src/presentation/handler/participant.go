package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	json_res "github.com/Doer-org/google-cloud-challenge-2022/utils/http/json"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/http/response"
	"github.com/go-chi/chi/v5"
)

type ParticipantHandler struct {
	uc usecase.IParticipantUsecase
}

func NewParticipantHandler(uc usecase.IParticipantUsecase) *ParticipantHandler {
	return &ParticipantHandler{
		uc: uc,
	}
}

func (h *ParticipantHandler) CreateNewParticipant(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	j := &json_res.ParticipantJson{}
	// TODO: bodyが空だった場合 "EOF"が入っている？
	if err := json.NewDecoder(r.Body).Decode(j); err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(http.StatusBadRequest,"StatusBadRequest",err),
			http.StatusBadRequest,
		)
		return
	}
	defer r.Body.Close()

	// TODO:commentがnullのときエラー
	participant, err := h.uc.CreateNewParticipant(
		r.Context(),
		j.Name,
		j.Comment.Body,
		idParam,
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
		json_res.EntityToJsonParticipant(participant),
		http.StatusOK,
	)
}
