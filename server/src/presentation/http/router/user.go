package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/handler"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/middleware"
)

func (r *Router) InitUser(userH *handler.User, m *middleware.Auth) {
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
}
