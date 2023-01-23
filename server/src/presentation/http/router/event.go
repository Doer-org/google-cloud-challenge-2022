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

func (r *Router) InitEvent(c *ent.Client) error {
	evenRepo := persistance.NewEvent(c)
	eventUC := usecase.NewEvent(evenRepo)
	eventH := handler.NewEvent(eventUC)

	// auth middleware
	authRepo := persistance.NewAuth(c)
	userRepo := persistance.NewUser(c)
	rg := google.NewClient("http://localhost:8080/auth/callback")
	authUC := usecase.NewAuth(authRepo, rg, userRepo)
	m := middleware.NewAuth(authUC)

	r.mux.Route("/events", func(r chi.Router) {
		r.Get("/{id}", eventH.GetEventById)
		r.Get("/{id}/admin", eventH.GetEventAdminById)
		r.Get("/{id}/comments", eventH.GetEventComments)
		r.Post("/{id}/participants", eventH.AddNewEventParticipant)
		r.Get("/{id}/users", eventH.GetEventUsers)

		// authentication required
		r.Route("/", func(r chi.Router) {
			r.Use(m.Authenticate)
			r.Post("/", eventH.CreateNewEvent)
			// TODO: sessionを持つユーザーであれば誰でも削除できるようになっているので、
			// そのeventのadminか確認する処理が必要。user apiも同様の処理がいる
			r.Delete("/{id}", eventH.DeleteEventById)
			r.Patch("/{id}", eventH.UpdateEventById)
			r.Patch("/{id}/state", eventH.ChangeEventStatusOfId)
		})
	})
	return nil
}
