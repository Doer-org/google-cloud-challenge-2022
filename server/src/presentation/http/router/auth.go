package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/google"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/persistance"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/handler"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/middleware"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/env"
)

func (r *Router) InitAuth(c *ent.Client) error {
	authRepo := persistance.NewAuth(c)
	userRepo := persistance.NewUser(c)

	callbackApi, err := env.GetEssentialEnv("GOOGLE_CALLBACK_API")
	if err != nil {
		return err
	}
	rg := google.NewClient(callbackApi)
	userUC := usecase.NewUser(userRepo)
	authUC := usecase.NewAuth(authRepo, rg, userRepo)
	h := handler.NewAuth(authUC, userUC)

	// auth middleware
	m := middleware.NewAuth(authUC)

	r.mux.Route("/auth", func(r chi.Router) {
		r.Get("/login", h.Login)
		r.Get("/callback", h.Callback)

		// authentication required
		r.Route("/", func(r chi.Router) {
			r.Use(m.Authenticate)
			r.Get("/validate", h.Validate)
			r.Get("/user", h.User)
		})
	})
	return nil
}
