package handler

import (
	"net/http"

	res "github.com/Doer-org/google-cloud-challenge-2022/presentation/http/response"
)

type IHealth interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type Health struct{}

func NewHealth() IHealth {
	return &Health{}
}

func (h *Health) Ping(w http.ResponseWriter, r *http.Request) {
	res.WriteJson(w, res.New200SuccessJson("pong!"), http.StatusOK)
}
