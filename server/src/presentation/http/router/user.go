package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/persistance"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/handler"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
)

func (r *Router) InitUser(c *ent.Client) {
	repo := persistance.NewUserRepository(c)
	uc   := usecase.NewUserUsecase(repo)
	h    := handler.NewUserHandler(uc)

	r.mux.Route("/users", func(r chi.Router) {
		r.Post("/", h.CreateNewUser)
		r.Get("/{id}", h.GetUserById)
		r.Delete("/{id}", h.DeleteUserById)
		r.Patch("/{id}", h.UpdateUserById)
		r.Get("/{id}/events", h.GetUserEvents)
		r.Get("/", h.GetUserByMail)
	})
}
