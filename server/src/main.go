package main

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/config"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/router"

	_ "github.com/lib/pq"
)

func main() {
	err := config.GetEnvAll()
	if err != nil {
		panic(fmt.Sprintf("error: GetEnvAll: %v", err))
	}
	client, err := ent.Open("postgres", config.POSTGRES_URL)
	if err != nil {
		panic(fmt.Sprintf("error: ent.Open: %v", err))
	}
	defer client.Close()

	// TODO: migrationsがうまくいかない
	if err := client.Schema.Create(context.Background()); err != nil {
		panic(fmt.Sprintf("error: Schema.Create: %v", err))
	}
	r, err := router.NewDefaultChiRouter(
		config.PORT,
		client,
	)
	if err != nil {
		panic(fmt.Sprint("error: NewDefaultChiRouter: %w", err))
	}
	if err := r.Serve(); err != nil {
		panic(fmt.Sprint("error: Serve: %w", err))
	}
}
