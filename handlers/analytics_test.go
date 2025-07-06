package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAnalyticsHandler_NoWallet(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/analytics", AnalyticsHandler)

	req, _ := http.NewRequest("GET", "/analytics", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestAnalyticsHandler_NoStrategies(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/analytics", AnalyticsHandlerWithDeps(
		func(wallet string, page, perPage int, status string) ([]interface{}, error) {
			return []interface{}{}, nil // no strategies
		},
		func(chainId int, protocol, poolAddress string) (interface{}, error) {
			return nil, nil
		},
		func(apiKey, prompt string) (string, error) {
			return "", nil
		},
		func() string { return "dummy" },
	))

	req, _ := http.NewRequest("GET", "/analytics?wallet=0xabc", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	if want := "No strategies found"; !contains(w.Body.String(), want) {
		t.Errorf("expected body to contain %q, got %s", want, w.Body.String())
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || (len(s) > len(substr) && contains(s[1:], substr))) || (len(s) >= len(substr) && s[:len(substr)] == substr)
}

// Add more tests for other scenarios as needed
