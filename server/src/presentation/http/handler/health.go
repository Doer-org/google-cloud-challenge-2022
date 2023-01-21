package handler

import (
	"net/http"

	res "github.com/Doer-org/google-cloud-challenge-2022/presentation/http/response"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Ping(w http.ResponseWriter, r *http.Request) {
	res.WriteJson(w, res.New200SuccessJson("pong"), http.StatusOK)
}
