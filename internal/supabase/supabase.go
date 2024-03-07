package supabase

import (
	"fmt"
	"os"

	"github.com/nedpals/supabase-go"
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

	return supabase.CreateClient(url, key), nil
}