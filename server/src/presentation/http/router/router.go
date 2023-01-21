package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	mymiddleware "github.com/Doer-org/google-cloud-challenge-2022/presentation/http/middleware"
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

func NewDefaultRouter(port string, c *ent.Client) *Router {
	r := NewRouter(port)
	// middleware
	r.SetMiddleware()
	// init all router
	r.InitHealth()
	r.InitAuth(c)
	r.InitEvent(c)
	r.InitUser(c)
	return r
}

func (r *Router) SetMiddleware() {
	// logger
	r.mux.Use(middleware.Logger)

	// recover
	r.mux.Use(middleware.Recoverer)

	// cors
	r.mux.Use(mymiddleware.Cors)

	// content type json
	r.mux.Use(mymiddleware.ContentTypeJson)

	//TODO: middlewareにusecaseを渡してもいいのか...
	// setAuthMiddleware(r, authUC)

	// auth
	// TODO: setAuth?

}

//TODO: uc?or u

func (r *Router) Serve() error {
	return http.ListenAndServe(
		fmt.Sprintf(":%s", r.port),
		r.mux,
	)
}