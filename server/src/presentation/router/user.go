package router

import (
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/persistance"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/handler"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	"github.com/go-chi/chi/v5"
)

func initUserRouter(r *chi.Mux, c *ent.Client) {
	repo := persistance.NewUserRepository(c)
	uc := usecase.NewUserUsecase(repo)
	h := handler.NewUserHandler(uc)

	r.Route("/user", func(r chi.Router) {
		r.Post("/", h.Create)
		r.Get("/{mail}", h.GetByMail)
	})
}
