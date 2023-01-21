package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/google"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/persistance"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/handler"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
)

func (r *Router) InitAuth(c *ent.Client) {
	authRepo := persistance.NewAuth(c)
	userRepo := persistance.NewUser(c)
	//TODO: 環境変数にすべき
	rg := google.NewClient("http://localhost:8080/api/callback")
	//TODO: 順番が気になる
	uc := usecase.NewAuth(authRepo, rg, userRepo)
	//TODO: frontendURLが空?
	h := handler.NewAuth(uc, "")
	//TODO: /apiにする必要ある? /authとか？
	r.mux.Route("/api", func(r chi.Router) {
		r.Get("/login", h.Login)
		r.Get("/callback", h.Callback)
	})
}
