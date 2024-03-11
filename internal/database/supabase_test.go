package database

import (
	"testing"

	"github.com/joho/godotenv"
)

func Test_SupabaseConnection(t *testing.T) {
	t.Run("Test loading env file", func(t *testing.T) {
		if err := godotenv.Load("../../.env"); err != nil {
			t.Fatalf("Unable to load env file: %s", err.Error())
		} 
	})

	t.Run("Test connection to supabase postgreSQL database", func(t *testing.T) {
		if err := godotenv.Load("../../.env"); err != nil {
			t.Fatalf("Unable to load env file: %s", err.Error())
		} 

		if _, err := CreateSupabaseClient(); err != nil {
			t.Fatalf("Unable to establish connection to database: %s", err.Error())
		}
	})
}