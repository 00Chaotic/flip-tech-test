package http

import (
	"encoding/json"
	"log"
	"net/http"
)

// SendJSONResponse is a utility function to send an HTTP JSON respnose with the
// specified struct and status code.
func SendJSONResponse(w http.ResponseWriter, res interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println("failed to encode response", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}
