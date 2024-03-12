package gemini

import (
	"encoding/base64"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func Test_MakeGeminiRequest(t *testing.T) {
	t.Run("Test making a request to Gemini", func(t *testing.T) {
		if err := godotenv.Load("../../.env"); err != nil {
			t.Fatalf("Unable to load env file: %s", err.Error())
		}

		encodedMSE, err := os.ReadFile("EncodedMSE.txt")
        if err != nil {
            t.Fatalf("Unable to read EncodedMSE.txt: %s", err.Error())
        }

		decodedImg, err := base64.StdEncoding.DecodeString(string(encodedMSE))
		if err != nil {
			t.Fatalf("decode error: %s", err.Error())
		}

		content, err := MakeGeminiRequest("png", "What equation is this and what does it do?", decodedImg)

		if err != nil {
			t.Fatalf("Unable to make request to Gemini: %s", err.Error())
		}

		t.Log(content)
	})
}