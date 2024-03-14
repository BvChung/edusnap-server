package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ocr/internal/database"
	"github.com/ocr/internal/env"
	"github.com/ocr/internal/handlers"
	"github.com/ocr/internal/middleware"
	"github.com/ocr/internal/vertexai"
	"github.com/rs/cors"
)

func main() {
	if err := env.LoadEnvVariables(".env"); err != nil {
		log.Fatal(err.Error())
	}

	supabaseClient, err := database.CreateSupabaseClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	vertexClient, err := vertexai.CreateVertexClient()
	if err != nil {
		log.Fatal(err.Error())
	}

	mux := http.NewServeMux()
	mux.Handle("POST /api/login", http.HandlerFunc(handlers.LoginHandler))
	mux.Handle("POST /api/register", http.HandlerFunc(handlers.RegisterHandler))
	mux.Handle("/api/messages", handlers.NewMessagesHandler(supabaseClient, vertexClient))
	mux.Handle("/api/s", handlers.NewStudentHandler(supabaseClient, vertexClient))

	mux.Handle("/", http.NotFoundHandler())

	s := &http.Server{
		Addr:           ":8080",
		Handler:        cors.Default().Handler(middleware.LoggingMiddleware(mux)),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server listening on Port 8080. Live at http://localhost:8080")

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
