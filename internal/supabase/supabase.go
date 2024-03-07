package supabase

import (
	"fmt"
	"os"

	"github.com/nedpals/supabase-go"
)

func CreateClient() (*supabase.Client, error) {
	url := os.Getenv("SUPABASE_URL")
	if url == "" {
		return nil, fmt.Errorf("supabase url required")
	}

	key := os.Getenv("SUPABASE_KEY")
	if key == "" {
		return nil, fmt.Errorf("supabase api key required")
	}

	client := supabase.CreateClient(url, key)

	return client, nil
}