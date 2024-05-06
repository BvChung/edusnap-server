package main

import (
	"net/http"

	"cloud.google.com/go/vertexai/genai"
	"github.com/ocr/cmd/handlers"
	"github.com/supabase-community/supabase-go"
)

func addRoutes(mux *http.ServeMux, supabaseClient *supabase.Client, vertexClient *genai.Client) {
	mux.Handle("/api/login", handlers.NewLoginHandler(supabaseClient))
	mux.Handle("/api/register", handlers.NewRegisterHandler(supabaseClient))
	mux.Handle("/api/messages", handlers.NewMessagesHandler(supabaseClient, vertexClient))
	mux.Handle("/", http.NotFoundHandler())
}
