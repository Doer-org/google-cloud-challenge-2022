package router

import (
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/persistance"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/handler"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	"github.com/go-chi/chi/v5"
)

func initParticipantHandler(r *chi.Mux, c *ent.Client) {
	repo := persistance.NewParticipantRepository(c)
	uc := usecase.NewParticipantUsecase(repo)
	h := handler.NewParticipantHandler(uc)

	r.Route("/event/participant", func(r chi.Router) {
		r.Post("/{id}", h.CreateNewParticipant)
	})
}
