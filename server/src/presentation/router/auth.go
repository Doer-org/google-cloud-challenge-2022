package router

import (
	"github.com/Doer-org/google-cloud-challenge-2022/google"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/persistance"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/handler"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	"github.com/go-chi/chi/v5"
)

func initAuthRouter(r *chi.Mux, c *ent.Client) {
	authrepo := persistance.NewAuthRepository(c)
	userrepo := persistance.NewUserRepository(c)
	googlecli := google.NewClient("")

	uc := usecase.NewAuthUsecase(authrepo, googlecli, userrepo)
	h := handler.NewAuthHandler(uc, "")

	setAuthMiddleware(r, uc)

	r.Route("", func(r chi.Router) {
		r.Get("/login", h.Login)
		r.Get("/callback", h.Callback)
	})
}
