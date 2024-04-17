package decoderequest

import (
	"encoding/json"
	"net/http"
)

func DecodeRequestBody(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
