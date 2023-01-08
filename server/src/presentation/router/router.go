package router

import (
	"fmt"
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/persistance"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/handler"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/helper"

	"github.com/go-chi/chi/v5"
)

func InitRouter(c *ent.Client) {
	r := chi.NewRouter()

	// middleware
	setMiddleware(r)

	// repository
	userRepo  := persistance.NewUserRepository(c)
	eventRepo := persistance.NewEventRepository(c)

	// usecsae
	userUC  := usecase.NewUserUsecase(userRepo)
	eventUC := usecase.NewEventUsecae(eventRepo)

	healthH := handler.NewHealthHandler()
	userH   := handler.NewUserHandler(userUC,eventUC)
	eventH  := handler.NewEventHandler(eventUC)


	// health handler
	r.Get("/ping", healthH.Ping)

	// user handler
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userH.CreateNewUser)
		r.Get("/{id}", userH.GetUserById)
		r.Delete("/{id}",userH.DeleteUserById)
		r.Patch("/{id}",userH.UpdateUserById)
		r.Get("/{id}/events",userH.GetUserEvents)
		r.Get("/", userH.GetUserByMail)
	})

	// event handler
	r.Route("/event", func(r chi.Router) {
		r.Post("/", eventH.CreateNewEvent)
		r.Get("/{id}", eventH.GetEventById)
		r.Patch("/{id}", eventH.ChangeEventStatusOfId)
	})

	initParticipantHandler(r, c)

	http.ListenAndServe(
		fmt.Sprintf(":%s", helper.GetEnvOrDefault("PORT", "8080")),
		r,
	)
}
