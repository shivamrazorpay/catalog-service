package customerr

import (
	"encoding/json"
	"net/http"
)

// JSONErrorResponse is a helper function that sends an error response in JSON format.
func JSONErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
}
