package auth

import (
	"encoding/base64"
	"net/http"
	"service-catalog/boot"
	"service-catalog/customerr"
	"strings"
)

// BasicAuthMiddleware is a middleware that verifies Basic Auth credentials.
func BasicAuthMiddleware() func(http.Handler) http.Handler {
	username := boot.GlobalConfig.Auth.Username
	password := boot.GlobalConfig.Auth.Password

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				customerr.JSONErrorResponse(w, http.StatusUnauthorized, "Authorization header missing")
				return
			}

			// Check if the header starts with "Basic "
			if !strings.HasPrefix(authHeader, "Basic ") {
				customerr.JSONErrorResponse(w, http.StatusUnauthorized, "Invalid authorization header")
				return
			}

			// Decode the Base64 encoded credentials
			encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
			decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
			if err != nil {
				customerr.JSONErrorResponse(w, http.StatusUnauthorized, "Invalid credentials")
				return
			}

			// Split the decoded credentials into username and password
			creds := strings.SplitN(string(decodedCredentials), ":", 2)
			if len(creds) != 2 {
				customerr.JSONErrorResponse(w, http.StatusUnauthorized, "Invalid credentials")
				return
			}

			reqUsername, reqPassword := creds[0], creds[1]

			// Validate credentials
			if reqUsername != username || reqPassword != password {
				customerr.JSONErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			// If credentials are valid, proceed to the next handler
			next.ServeHTTP(w, r)
		})
	}
}
