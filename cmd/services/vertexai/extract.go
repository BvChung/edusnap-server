package vertexai

import (
	"fmt"
	"strings"
)

func ExtractBase64(s string) (b64 string, mimeType string, err error) {
	if len([]byte(s)) == 0 {
		return "", "", fmt.Errorf("invalid base64 encoded string")
	}

	before, encoded, found := strings.Cut(s, ";base64,")
	b64 = encoded

	if !found {
		return "", "", fmt.Errorf("invalid base64 encoded string")
	}

	_, ftype, found := strings.Cut(before, "image/")
	mimeType = ftype

	if !found {
		return "", "", fmt.Errorf("invalid base64 encoded string")
	}

	return b64, mimeType, nil
}
