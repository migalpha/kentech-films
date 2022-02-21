package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"runtime"
	"time"
)

// @title Kentech-Films
// @version 1.0.0
// @description This API provides endpoints to manage films.
// @description [Read me](https://github.com/migalpha/kentech-films)
// @schemes http
// @host localhost:8080
// @BasePath /api/films
func main() {
	// The first defered call is the last to be executed
	// os.Exit terminates the program
	defer os.Exit(0)

	var app application

	log.Print("Setting up dependecies")
	err := app.setupDependencies()
	if err != nil {
		log.Fatalf("Can't setup dependecies: %s", err.Error())
	}

	log.Print("Setting up server")
	srv := setupServer(app)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %s", err.Error())
	}

	// Handle graceful shutdowns via SIGINT
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Wait for requests to finish
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Graceful shutdown
	log.Println("Shutting down server gracefully")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Shutting down server error: %s\n", err.Error())
	}

	// Call all defered calls before closing server
	runtime.Goexit()
}
