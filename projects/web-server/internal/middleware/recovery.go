package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime/debug"
)

// Recovery captura panics y evita que el servidor se caiga
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf(
					"[PANIC] %s %s - Error: %v\n%s",
					r.Method,
					r.URL.Path,
					err,
					string(debug.Stack()),
				)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"error":   "Internal Server Error",
					"message": "An unexpected error occurred",
				})
			}
		}()

		next.ServeHTTP(w, r)
	})
}
