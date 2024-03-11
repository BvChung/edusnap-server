package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables(path string) error {
	if err := godotenv.Load(path); err != nil {
		return fmt.Errorf("unable to load .env file")
	}

	envKeys := []string{
		"SUPABASE_URL",
		"SUPABASE_KEY",
		"GOOGLE_APPLICATION_CREDENTIALS",
		"GOOGLE_CLOUD_PROJECT_ID",
		"GOOGLE_CLOUD_REGION",
		"GOOGLE_CLOUD_VERTEX_MODEL_NAME",
	}

	for _, key := range envKeys {
		if os.Getenv(key) == "" {
			return fmt.Errorf("missing env variable key or value %s", key)
		}
	}

	return nil
}