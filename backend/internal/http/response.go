package http

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, res interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println("failed to encode response", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}
