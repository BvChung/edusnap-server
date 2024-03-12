package gemini

import "testing"

func Test_ExtractString(t *testing.T) {
	data := []struct{
		Input string
		ExpectedB64 string
		ExpectedMimeType string
	}{
		{"", "iVBORw0KGgoAAAANSUhEUgAAAkw", "png"},
		{"data:image/png;iVBORw0KGgoAAAANSUhEUgAAAkw", "iVBORw0KGgoAAAANSUhEUgAAAkw", "png"},
		{"data:image;base64,iVBORw0KGgoAAAANSUhEUgAAAkw", "iVBORw0KGgoAAAANSUhEUgAAAkw", "png"},
		{"data:image/png;base64,", "iVBORw0KGgoAAAANSUhEUgAAAkw", "png"},
		{"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAkw", "iVBORw0KGgoAAAANSUhEUgAAAkw", "png"},
	}

	for _, d := range data {
		t.Run("Test extacting base64 encoded string", func(t *testing.T) {
			b64, mimeType, err := ExtractBase64(d.Input)
			
			if err != nil {
				t.Fatal(err.Error())
			}
	
			if d.ExpectedB64 != b64 || d.ExpectedMimeType != mimeType {
				t.Errorf("Expected: %s type: %s, got %s type: %s", d.ExpectedB64, d.ExpectedMimeType, b64, mimeType)
			}
		})
	}
}