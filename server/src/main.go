package main

import (
	"context"
	"fmt"

	"github.com/Doer-org/google-cloud-challenge-2022/config"
	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/http/router"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/env"

	_ "github.com/lib/pq"
)

// TODO: eventに制限時間を持たせる
func main() {
	dsn, err := config.DSN()
	if err != nil {
		panic(fmt.Sprintf("error: DSN: %v", err))
	}
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		panic(fmt.Sprintf("error: ent.Open: %v", err))
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		panic(fmt.Sprintf("error: Schema.Create: %v", err))
	}
	r, err := router.NewDefaultRouter(
		env.GetEnvOrDefault("PORT", "8080"),
		client,
	)
	if err != nil {
		panic(fmt.Sprint("error: NewDefaultRouter: %w", err))
	}
	if err := r.Serve(); err != nil {
		panic(fmt.Sprint("error: Serve: %w", err))
	}
}
