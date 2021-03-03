package api

import (
	"encoding/json"
	"net/http"
)

// response sends a JSON with status code
func response(w http.ResponseWriter, data map[string]interface{}, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	switch statusCode {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 201:
		w.WriteHeader(http.StatusCreated)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	case 404:
		w.WriteHeader(http.StatusNotFound)
	case 409:
		w.WriteHeader(http.StatusConflict)
	case 422:
		w.WriteHeader(http.StatusUnprocessableEntity)
	case 500:
		w.WriteHeader(http.StatusInternalServerError)
	case 503:
		w.WriteHeader(http.StatusServiceUnavailable)
	default:
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(data)
}
