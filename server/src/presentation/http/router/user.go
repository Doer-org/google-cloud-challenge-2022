package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/google"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/persistance"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/handler"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/middleware"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
)

func (r *Router) InitUser(c *ent.Client) error {
	userRepo := persistance.NewUser(c)
	userUC   := usecase.NewUser(userRepo)
	userH    := handler.NewUser(userUC)

	// auth middleware
	authRepo := persistance.NewAuth(c)
	rg       := google.NewClient("http://localhost:8080/auth/callback")
	authUC   := usecase.NewAuth(authRepo, rg, userRepo)
	m        := middleware.NewAuth(authUC)

	r.mux.Route("/users", func(r chi.Router) {
		// r.Post("/", userH.CreateNewUser) // /auth/loginからたたくのでコメントに
		r.Get("/{id}", userH.GetUserById)
		r.Get("/{id}/events", userH.GetUserEvents)
		r.Get("/", userH.GetUserByMail)

		// authentication required
		r.Route("/", func(r chi.Router) {
			r.Use(m.Authenticate)
			r.Delete("/{id}", userH.DeleteUserById)
			r.Patch("/{id}", userH.UpdateUserById)
		})
	})
	return nil
}
