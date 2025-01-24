package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extract API Key from headers
// Example:
// Authorization: ApiKey {key}
func GetAPIKey(headers http.Header) (string, error) {
	auth := headers.Get("Authorization")
	if auth == "" {
		return "", errors.New("no auth info found")
	}

	vals := strings.Split(auth, " ")
	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("auth info is malformed")
	}

	return vals[1], nil
}
