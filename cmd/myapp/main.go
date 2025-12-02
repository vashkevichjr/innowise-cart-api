package main

import (
	"context"
	"log"

	"github.com/vashkevichjr/innowise-cart-api/internal/app"
)

func main() {
	ctx := context.Background()
	myapp, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if err := myapp.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
