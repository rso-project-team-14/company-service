package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoints(t *testing.T) {
	tests := []struct {
		path         string
		expectedBody string
	}{
		{"/health/live", "OK\n"},
		{"/health/ready", "READY\n"},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			rec := httptest.NewRecorder()

			newHandler().ServeHTTP(rec, req)

			if rec.Code != http.StatusOK {
				t.Errorf("expected status 200, got %d", rec.Code)
			}

			if rec.Body.String() != tt.expectedBody {
				t.Errorf("expected %q, got %q",
					tt.expectedBody, rec.Body.String())
			}
		})
	}
}
