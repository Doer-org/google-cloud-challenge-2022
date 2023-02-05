package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Doer-org/google-cloud-challenge-2022/config"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/google"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/persistence"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/handler"
	authmiddleware "github.com/Doer-org/google-cloud-challenge-2022/presentation/http/middleware"
	mymiddleware "github.com/Doer-org/google-cloud-challenge-2022/presentation/http/middleware"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
)

type IRouter interface {
	Serve() error
	SetMiddlewares()
	InitHealth(healthH handler.IHealth)
	InitUser(userH handler.IUser, m authmiddleware.IAuth)
	InitEvent(eventH handler.IEvent, m authmiddleware.IAuth)
	InitAuth(eventH handler.IAuth, m authmiddleware.IAuth)
}

type ChiRouter struct {
	mux  *chi.Mux
	port string
}

func NewChiRouterImpl(port string) IRouter {
	return &ChiRouter{
		mux:  chi.NewRouter(),
		port: port,
	}
}

func NewDefaultChiRouter(port string, c *ent.Client) (IRouter, error) {
	r := NewChiRouterImpl(port)

	r.SetMiddlewares()

	rg := google.NewClient(config.GOOGLE_CALLBACK_API)

	userRepo := persistence.NewUser(c)
	authRepo := persistence.NewAuth(c)
	evenRepo := persistence.NewEvent(c)

	userUC := usecase.NewUser(userRepo)
	authUC := usecase.NewAuth(authRepo, rg, userRepo)
	eventUC := usecase.NewEvent(evenRepo)

	userH := handler.NewUser(userUC)
	authH := handler.NewAuth(authUC, userUC)
	eventH := handler.NewEvent(eventUC)
	healthH := handler.NewHealth()

	m := authmiddleware.NewAuth(authUC)

	r.InitHealth(healthH)
	r.InitUser(userH, m)
	r.InitEvent(eventH, m)
	r.InitAuth(authH, m)

	return r, nil
}

func (r *ChiRouter) SetMiddlewares() {
	r.setMiddlewares(
		middleware.Logger,
		middleware.Recoverer,
		mymiddleware.Cors,
		mymiddleware.ContentTypeJson,
	)
}

func (r *ChiRouter) setMiddlewares(middlewares ...func(next http.Handler) http.Handler) {
	for _, middleware := range middlewares {
		r.mux.Use(middleware)
	}
}

func (r *ChiRouter) Serve() error {
	return http.ListenAndServe(
		fmt.Sprintf(":%s", r.port),
		r.mux,
	)
}
