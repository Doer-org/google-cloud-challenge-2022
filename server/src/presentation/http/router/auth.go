package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/google"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/persistance"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/handler"
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
	uc := usecase.NewAuth(authRepo, rg, userRepo)
	//TODO: frontendURLが空?
	h := handler.NewAuth(uc, "")
	r.mux.Route("/auth", func(r chi.Router) {
		r.Get("/login", h.Login)
		r.Get("/callback", h.Callback)
	})
	return nil
}
