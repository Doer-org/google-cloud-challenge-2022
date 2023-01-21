package router

import (
	"fmt"
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/google"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/persistance"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/handler"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/helper"

	"github.com/go-chi/chi/v5"
)

func InitRouter(c *ent.Client) {
	r := chi.NewRouter()

	// middleware
	setMiddleware(r)

	// repository
	userRepo := persistance.NewUserRepository(c)
	eventRepo := persistance.NewEventRepository(c)
	authRepo := persistance.NewAuthRepository(c)

	googlecli := google.NewClient("http://localhost:8080/api/callback")

	// usecsae
	userUC := usecase.NewUserUsecase(userRepo)
	eventUC := usecase.NewEventUsecae(eventRepo)
	authUC := usecase.NewAuthUsecase(authRepo, googlecli, userRepo)

	healthH := handler.NewHealthHandler()
	userH := handler.NewUserHandler(userUC)
	eventH := handler.NewEventHandler(eventUC)
	authH := handler.NewAuthHandler(authUC, "")

	// health handler
	r.Get("/ping", healthH.Ping)

	// auth handler
	r.Route("/api", func(r chi.Router) {
		r.Get("/login", authH.Login)
		r.Get("/callback", authH.Callback)
	})

	// user handler
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userH.CreateNewUser)
		r.Get("/{id}", userH.GetUserById)
		r.Delete("/{id}", userH.DeleteUserById)
		r.Patch("/{id}", userH.UpdateUserById)
		r.Get("/{id}/events", userH.GetUserEvents)
		r.Get("/", userH.GetUserByMail)
	})

	// event handler
	r.Route("/events", func(r chi.Router) {
		r.Post("/", eventH.CreateNewEvent)
		r.Get("/{id}", eventH.GetEventById)
		r.Delete("/{id}", eventH.DeleteEventById)
		r.Patch("/{id}", eventH.UpdateEventById)
		r.Get("/{id}/admin", eventH.GetEventAdminById)
		r.Get("/{id}/comments", eventH.GetEventComments)
		r.Post("/{id}/participants", eventH.AddNewEventParticipant)
		r.Patch("/{id}/state", eventH.ChangeEventStatusOfId)
		r.Get("/{id}/users", eventH.GetEventUsers)
	})

	// comment handler
	// r.Route("/comments",func(r chi.Router) {
	// 	r.Post("/",commentH.CreateNewComment)
	// })

	// setAuthMiddleware(r, authUC)
	//TODO: 消す
	r.Get("/pong", healthH.Pong)

	// TODO: errハンドリング
	http.ListenAndServe(
		fmt.Sprintf(":%s", helper.GetEnvOrDefault("PORT", "8080")),
		r,
	)
}
