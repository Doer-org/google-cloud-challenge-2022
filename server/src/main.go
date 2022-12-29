package main

import (
	"fmt"
	"net/http"

	"context"
	"log"

	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"github.com/Doer-org/google-cloud-challenge-2022/presentation/router"

	_ "github.com/lib/pq"
)


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {

    client, err := ent.Open("postgres","host=google-cloud-challenge-2022-db port=5432 user=hoge dbname=hoge password=hoge sslmode=disable")
    if err != nil {
        log.Fatalf("failed opening connection to postgres: %v", err)
    }
    defer client.Close()

    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

	router.InitRouter(client)
}
