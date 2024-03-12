package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ocr/internal/env"
	"github.com/ocr/internal/handlers"
	"github.com/rs/cors"
)

func main() {
	if err := env.LoadEnvVariables(".env"); err != nil {
		log.Fatal(err.Error())
	}
	
	mux := http.NewServeMux()

	mux.Handle("POST /api/login", http.HandlerFunc(handlers.LoginHandler))
	mux.Handle("POST /api/register", http.HandlerFunc(handlers.RegisterHandler))
	mux.Handle("POST /api/message", http.HandlerFunc(handlers.MessageHandler))
	mux.Handle("GET /api/s", http.HandlerFunc(handlers.StudentHandler))

	mux.HandleFunc("GET /profile/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		// forms := r.Form

		search := r.URL.Query().Get("search")
		fmt.Println(search)
		fmt.Fprintf(w, "Parameter: %s", id)
	})

	mux.Handle("/", http.NotFoundHandler())

	s := &http.Server{
		Addr:           ":8080",
		Handler:        cors.Default().Handler(mux),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server listening on Port 8080. Live at http://localhost:8080")

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
