package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/Doer-org/google-cloud-challenge-2022/presentation/handler"
)

func initHealthRouer(r *chi.Mux){
	h := handler.NewHealthHandler()
	r.Get("/ping",h.Ping)
}
