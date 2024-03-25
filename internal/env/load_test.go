package env

import (
	"testing"
)

func Test_LoadEnv(t *testing.T) {
	envKeys := []string{
		"SUPABASE_URL",
		"SUPABASE_KEY",
		"GOOGLE_APPLICATION_CREDENTIALS",
		"GOOGLE_CLOUD_PROJECT_ID",
		"GOOGLE_CLOUD_REGION",
		"GOOGLE_CLOUD_VERTEX_MODEL_NAME",
	}

	for _, key := range envKeys{
		t.Run("Test missing env variables error", func(t *testing.T) {
			t.Setenv(key, "")

			if err := LoadEnvVariables("../../.env"); err != nil {
				t.Fatal(err.Error())
			}
		})
	}

	t.Run("Test loading env variables", func(t *testing.T) {
		if err := LoadEnvVariables("../../.env"); err != nil {
			t.Fatal(err.Error())
		}
	})
}
