package router

import (
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/handler"
)

func (r *ChiRouter) InitHealth(healthH handler.IHealth) {
	r.mux.Get("/ping", healthH.Ping)
}
