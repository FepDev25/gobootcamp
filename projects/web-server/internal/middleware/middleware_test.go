package middleware

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestRecovery(t *testing.T) {
	handler := Recovery(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("something went wrong")
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, rr.Code)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "Internal Server Error") {
		t.Errorf("expected error message in body, got %s", body)
	}
}

func TestRecoveryNormalRequest(t *testing.T) {
	handler := Recovery(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	if rr.Body.String() != "success" {
		t.Errorf("expected body 'success', got %s", rr.Body.String())
	}
}

func TestCORS(t *testing.T) {
	tests := []struct {
		name           string
		allowedOrigins []string
		requestOrigin  string
		expectAllowed  bool
	}{
		{
			name:           "allow specific origin",
			allowedOrigins: []string{"http://localhost:3000"},
			requestOrigin:  "http://localhost:3000",
			expectAllowed:  true,
		},
		{
			name:           "deny unknown origin",
			allowedOrigins: []string{"http://localhost:3000"},
			requestOrigin:  "http://evil.com",
			expectAllowed:  false,
		},
		{
			name:           "allow all with wildcard",
			allowedOrigins: []string{"*"},
			requestOrigin:  "http://any-origin.com",
			expectAllowed:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := CORS(tt.allowedOrigins)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			req.Header.Set("Origin", tt.requestOrigin)
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			origin := rr.Header().Get("Access-Control-Allow-Origin")
			if tt.expectAllowed && origin == "" {
				t.Error("expected CORS headers to be set")
			}
			if !tt.expectAllowed && origin != "" {
				t.Error("expected no CORS headers for denied origin")
			}
		})
	}
}

func TestCORSOptions(t *testing.T) {
	handler := CORS([]string{"*"})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("should not reach here"))
	}))

	req := httptest.NewRequest(http.MethodOptions, "/test", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d for OPTIONS, got %d", http.StatusOK, rr.Code)
	}

	// Verificar que no se ejecutó el handler
	if rr.Body.String() == "should not reach here" {
		t.Error("handler should not be called for OPTIONS request")
	}
}

func TestRequestID(t *testing.T) {
	handler := RequestID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	requestID := rr.Header().Get("X-Request-ID")
	if requestID == "" {
		t.Error("expected X-Request-ID header to be set")
	}

	// Verificar que el ID es un UUID válido (36 caracteres)
	if len(requestID) != 36 {
		t.Errorf("expected UUID of length 36, got %d", len(requestID))
	}
}

func TestRequestIDPreservesExisting(t *testing.T) {
	handler := RequestID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("X-Request-ID", "existing-id-123")
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	requestID := rr.Header().Get("X-Request-ID")
	if requestID != "existing-id-123" {
		t.Errorf("expected existing ID, got %s", requestID)
	}
}

func TestGetRequestID(t *testing.T) {
	handler := RequestID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := GetRequestID(r.Context())
		if id == "" {
			t.Error("expected to get request ID from context")
		}
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)
}

func TestTimeout(t *testing.T) {
	handler := Timeout(50 * time.Millisecond)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusGatewayTimeout {
		t.Errorf("expected status %d, got %d", http.StatusGatewayTimeout, rr.Code)
	}
}

func TestTimeoutSuccess(t *testing.T) {
	handler := Timeout(100 * time.Millisecond)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestLogger(t *testing.T) {
	handler := Logger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))

	req := httptest.NewRequest(http.MethodPost, "/api/test", nil)
	req.Header.Set("User-Agent", "test-agent")
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, rr.Code)
	}
}

func TestChainedMiddleware(t *testing.T) {
	// Test que verifica que los middleware funcionan en cadena
	var executionOrder []string

	handler := Recovery(
		RequestID(
			Logger(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					executionOrder = append(executionOrder, "handler")
					w.WriteHeader(http.StatusOK)
				}),
			),
		),
	)

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	if len(executionOrder) != 1 || executionOrder[0] != "handler" {
		t.Errorf("handler was not executed properly: %v", executionOrder)
	}
}
