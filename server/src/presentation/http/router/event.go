package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/handler"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/middleware"
)

func (r *ChiRouter) initEvent(eventH handler.IEvent, m middleware.IAuth) {
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
			// そのeventのadminか確認する処理が必要。user apiも同様の処理がいる
			r.Delete("/{id}", eventH.DeleteEventById)
			r.Patch("/{id}", eventH.UpdateEventById)
			r.Patch("/{id}/state", eventH.ChangeEventStatusOfId)
		})
	})
}
