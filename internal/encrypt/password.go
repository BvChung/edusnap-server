package encrypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	pw := []byte(password)

	if len(pw) == 0 {
		return nil, fmt.Errorf("password cannot be an empty string")
	}

	hash, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("unable to hash password, %w", err)
	}

	return hash, nil
}
