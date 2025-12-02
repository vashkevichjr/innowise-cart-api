package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vashkevichjr/innowise-cart-api/internal/app"
)

func main() {
	ctx := context.Background()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	myapp, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal(err)
	}

	errChan := make(chan error, 1)
	go func() {
		if err := myapp.Run(ctx); err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		log.Fatalf("Error starting: %v", err)
	case <-quit:
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := myapp.Close(ctx); err != nil {
			log.Fatalf("Error shutting down: %v", err)
		}

		log.Println("Server stopped")
	}
}
