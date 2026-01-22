package handlers

import (
	"encoding/json"
	"net/http"
	"web-server/internal/database"
)

func HealthCheck(s database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		health := s.Health()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(health)
	}
}
