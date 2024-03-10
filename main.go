package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/ocr/internal/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load env")
	}

	fmt.Println("Server listening on Port 8080. Live at http://localhost:8080")
	mux := http.NewServeMux()

	// mux.Handle("GET /image", http.HandlerFunc(handlers.ImageHandler))
	mux.Handle("POST /login", http.HandlerFunc(handlers.LoginHandler))
	mux.Handle("POST /register", http.HandlerFunc(handlers.RegisterHandler))
	mux.Handle("POST /message", http.HandlerFunc(handlers.MessageHandler))
	mux.Handle("GET /s", http.HandlerFunc(handlers.StudentHandler))

	mux.HandleFunc("GET /profile/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		// forms := r.Form

		search := r.URL.Query().Get("search")
		fmt.Println(search)
		fmt.Fprintf(w, "Parameter: %s", id)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
