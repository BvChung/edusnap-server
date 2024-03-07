package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/ocr/internal/handlers"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Unable to load env")
	}

	fmt.Println("Server listening on Port 8080. Live at http://localhost:8080")
	mux := http.NewServeMux()

	// supabaseUrl := os.Getenv("SUPABASE_URL")
	// supabaseKey := os.Getenv("SUPABASE_KEY")
	

	mux.Handle("GET /image", http.HandlerFunc(handlers.ImageHandler))
	mux.Handle("GET /student", http.HandlerFunc(handlers.StudentHandler))

	mux.Handle("POST /login", http.HandlerFunc(handlers.LoginHandler))

	mux.Handle("POST /register", http.HandlerFunc(handlers.RegisterHandler))

	mux.HandleFunc("GET /profile/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		// forms := r.Form

		search := r.URL.Query().Get("search")
		fmt.Println(search)
		fmt.Fprintf(w, "Parameter: %s", id)
	})

	// mux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
	//      fmt.Fprintf(w, "Login user")
	// })

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatalf(err.Error())
	}
}
