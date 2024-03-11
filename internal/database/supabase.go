package database

import (
	"fmt"
	"os"

	"github.com/supabase-community/supabase-go"
)

func CreateSupabaseClient() (*supabase.Client, error) {
	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_KEY")

	client, err := supabase.NewClient(url, key, nil); if err != nil {
		return nil, fmt.Errorf("unable to connect to supabase database: %w", err)
	}

	return client, nil
}