package router

import (
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/handler"
)

func (r *Router) InitHealth(healthH *handler.Health) {
	r.mux.Get("/ping", healthH.Ping)
}
