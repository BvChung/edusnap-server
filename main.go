package main

import (
	supabase "github.com/nedpals/supabase-go"
    "fmt"
    "context"
	"log"
	"net/http"

	user "github.com/ocr/api/user"
)

func main() {
	fmt.Println("Server listening on Port 8080. Live at http://localhost:8080")
	mux := http.NewServeMux()

	mux.Handle("GET /login", http.HandlerFunc(user.CountWrapper(user.UserHandler)))
	mux.Handle("POST /login", http.HandlerFunc(user.LoginHandler))

	mux.HandleFunc("GET /profile/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

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
