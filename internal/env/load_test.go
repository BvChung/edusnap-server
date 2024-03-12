package env

import (
	"testing"
)

func Test_LoadEnv(t *testing.T) {
	t.Run("Test loading env variables", func(t *testing.T) {
		if err := LoadEnvVariables("../../.env"); err != nil {
			t.Fatal(err.Error())
		}
	})
}
