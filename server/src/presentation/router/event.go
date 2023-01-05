package router

import (
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/persistance"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/handler"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	"github.com/go-chi/chi/v5"
)

func initEventHandler(r *chi.Mux, c *ent.Client) {
	repo := persistance.NewEventRepository(c)
	uc := usecase.NewEventUsecae(repo)
	h := handler.NewEventHandler(uc)

	r.Route("/event", func(r chi.Router) {
		r.Post("/", h.Create)
		r.Get("/{id}", h.GetById)
	})
}
