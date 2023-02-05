package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/google"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/persistance"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/handler"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/env"
	mymiddleware "github.com/Doer-org/google-cloud-challenge-2022/presentation/http/middleware"
	authmiddleware"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/middleware"
	
)

type Router struct {
	mux  *chi.Mux
	port string
}

func NewRouter(port string) *Router {
	return &Router{
		mux:  chi.NewRouter(),
		port: port,
	}
}

func NewDefaultRouter(port string, c *ent.Client) (*Router, error) {
	r := NewRouter(port)

	r.SetMiddlewares()

	callbackApi, _ := env.GetEssentialEnv("GOOGLE_CALLBACK_API")
	rg := google.NewClient(callbackApi)

	userRepo := persistance.NewUser(c)
	authRepo := persistance.NewAuth(c)
	evenRepo := persistance.NewEvent(c)

	userUC := usecase.NewUser(userRepo)
	authUC := usecase.NewAuth(authRepo, rg, userRepo)
	eventUC := usecase.NewEvent(evenRepo)

	userH := handler.NewUser(userUC)
	authH := handler.NewAuth(authUC, userUC)
	eventH := handler.NewEvent(eventUC)
	healthH := handler.NewHealth()

	m := authmiddleware.NewAuth(authUC)

	r.InitHealth(healthH)
	r.InitUser(userH,m)
	r.InitEvent(eventH,m)
	r.InitAuth(authH,m)

	return r, nil
}

func (r *Router) SetMiddlewares() {
	r.setMiddlewares(
		middleware.Logger,
		middleware.Recoverer,
		mymiddleware.Cors,
		mymiddleware.ContentTypeJson,
	)

}

func (r *Router) setMiddlewares(middlewares ...func(next http.Handler) http.Handler) {
	for _,middleware := range middlewares {
		r.mux.Use(middleware)
	}
}

func (r *Router) Serve() error {
	return http.ListenAndServe(
		fmt.Sprintf(":%s", r.port),
		r.mux,
	)
}
