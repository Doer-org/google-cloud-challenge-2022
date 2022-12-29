package router

import (
	"fmt"
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/helper"

	"github.com/go-chi/chi/v5"
)

func InitRouter(c *ent.Client) {
	r := chi.NewRouter()

	setMiddleware(r)
	initHealthRouer(r)
	initUserRouter(r,c)

	r.Get("/",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok!"))
	})

	http.ListenAndServe(
		fmt.Sprintf(":%s",helper.GetEnvOrDefault("PORT","8080")),
		r,
	)
}
