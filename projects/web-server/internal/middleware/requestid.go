package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// contextKey es un tipo para evitar colisiones en context keys
type contextKey string

const requestIDKey contextKey = "requestID"

// RequestID agrega un identificador único a cada request para tracing
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Agregar al header de respuesta
		w.Header().Set("X-Request-ID", requestID)

		// Agregar al context
		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetRequestID obtiene el request ID del contexto
func GetRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(requestIDKey).(string); ok {
		return id
	}
	return ""
}
