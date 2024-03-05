package main

import (
	"fmt"
	"net/http"
	u "github.com/ocr/user"
)

func main() {
	fmt.Println("Server live at port 8080")
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello 🐳🐳")
	})

	mux.Handle("GET /login", http.HandlerFunc(u.HandleLogin))

	// mux.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Login user")
	// })

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		fmt.Println(err.Error())
	}
}