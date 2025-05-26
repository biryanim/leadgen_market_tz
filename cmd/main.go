package main

import (
	"context"
	"github.com/biryanim/leadgenmarket_tz/internal/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("failed to run: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	log.Printf("Received signal: %v. Gracefully shut down", sig)

}
