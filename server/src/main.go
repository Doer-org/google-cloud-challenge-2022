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
	dsn, err := config.DSN()
	if err != nil {
		panic(fmt.Sprintf("failed to get DSN : %v", err))
	}
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		panic(fmt.Sprintf("failed opening connection to postgres: %v", err))
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		panic(fmt.Sprintf("failed creating schema resources: %v", err))
	}

	router.InitRouter(client)
}
