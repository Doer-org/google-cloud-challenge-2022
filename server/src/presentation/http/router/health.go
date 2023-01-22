package router

import (
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/handler"
)

func (r *Router) InitHealth() error {
	h := handler.NewHealth()
	r.mux.Get("/ping", h.Ping)
	return nil
}
