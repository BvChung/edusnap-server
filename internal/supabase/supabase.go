package supabase

import (
	"fmt"
	"os"

	"github.com/supabase-community/supabase-go"
)

func CreateClient() (*supabase.Client, error) {
	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_KEY")
	

	if url == "" {
		return nil, fmt.Errorf("supabase url required")
	}

	if key == "" {
		return nil, fmt.Errorf("supabase api key required")
	}

	client, err := supabase.NewClient(url, key, nil); if err != nil {
		return nil, fmt.Errorf("unable to connect to supabase database: %w", err)
	}

	return client, nil
}