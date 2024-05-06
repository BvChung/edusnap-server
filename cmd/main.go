package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ocr/cmd/database"
	"github.com/ocr/cmd/env"
	"github.com/ocr/cmd/middleware"
	"github.com/ocr/cmd/services/vertexai"
	"github.com/rs/cors"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}


func run() error{
	if err := env.LoadEnvVariables(".env"); err != nil {
		return err
	}

	supabaseClient, err := database.CreateSupabaseClient()
	if err != nil {
		return err
	}

	vertexClient, err := vertexai.CreateVertexClient()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	addRoutes(mux, supabaseClient, vertexClient)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        cors.Default().Handler(middleware.LoggingMiddleware(mux)),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server listening on Port 8080. Live at http://localhost:8080")

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}